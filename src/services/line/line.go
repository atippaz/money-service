package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	auth "money-service/src/services/auth"
	expense "money-service/src/services/expense"
	tag "money-service/src/services/tag"
	user "money-service/src/services/user"

	"time"

	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func NewLineService(authService auth.AuthService, userService user.UserService, expenseService expense.ExpenseService, tagService tag.TagService) LineService {
	return &lineService{userService, authService, expenseService, tagService}
}
func (s *lineService) login(username string, password string) (*user.UserLoginInfo, error) {
	_, errLogin := s.authService.Login(username, password)
	if errLogin != nil {
		return nil, errLogin
	}
	user, err := s.userService.GetLoginDataByCredential(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *lineService) addIncome(username string, password string) (*user.UserLoginInfo, error) {
	_, errLogin := s.authService.Login(username, password)
	if errLogin != nil {
		return nil, errLogin
	}
	user, err := s.userService.GetLoginDataByCredential(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type SummaryTransection struct {
	Type  string          `json:type`
	Value decimal.Decimal `json:value`
}
type Tag struct {
	tag  string
	text string
}

func (s *lineService) summaryExpense(userId string) *[]SummaryTransection {
	now := time.Now()

	month := now.Month()
	year := now.Year()

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	endOfMonth := time.Date(year, time.Month(month)+1, 0, 23, 59, 59, 999999999, time.UTC)

	res, err := s.expenseService.GetSummary(uuid.MustParse(userId), &startOfMonth, &endOfMonth)
	tags := make(map[string]string)

	ct, _ := s.customTagService.GetTagsByUser(uuid.MustParse(userId), true)

	for _, result := range *ct {
		tags[result.TagId.String()] = result.NameTh
	}

	results := []SummaryTransection{}
	if err != nil {
		return &results
	}
	for _, result := range *res {
		results = append(results, SummaryTransection{
			Type:  tags[result.TagId.String()],
			Value: result.Value,
		})
	}
	return &results
}

type UserIdMapping struct {
	id    string
	token string
}

var dataLogin = make(map[string]UserIdMapping)

func (s *lineService) CommandManager(payload *LineMessage) {
	fmt.Println(payload.Events[0].Message.Text)
	splitText := strings.Split(payload.Events[0].Message.Text, "\n")
	command := splitText[0]
	text := Text{
		Type: "text",
		Text: "คุณสั่ง: " + command,
	}
	replytoken := payload.Events[0].ReplyToken
	userToken := payload.Events[0].Source.UserID
	message := ReplyMessage{
		ReplyToken: replytoken,
		Messages: []Text{
			text,
		},
	}
	if command == "login" {
		user, err := s.login(splitText[1], splitText[2])
		if err != nil {
			message.Messages[0].Text = "เกิดข้อผิดพลาด รหัสผ่านผิด \n" + err.Error()
			replyMessageLine(message)
			return
		}

		if _, exists := dataLogin[userToken]; exists {
			dataLogin[userToken] = UserIdMapping{
				id:    user.UserId.String(),
				token: userToken,
			}
		} else {
			dataLogin[userToken] = UserIdMapping{
				id:    user.UserId.String(),
				token: userToken,
			}
		}
		message.Messages[0].Text = "login สำเร็จ \nสวัสดีคุณ " + user.Email
	}
	userId := ""
	fmt.Println(dataLogin)

	if exist, exists := dataLogin[userToken]; exists {
		userId = exist.id
	}
	fmt.Println(userId)
	if command == "expense summary" {

		val := s.summaryExpense(userId)
		var sb strings.Builder
		for _, summary := range *val {
			sb.WriteString(fmt.Sprintf("ประเภท: %s\n", summary.Type))
			sb.WriteString(fmt.Sprintf("จำนวน: %s\n", summary.Value.String()))
			sb.WriteString("----------\n")
		}

		if sb.Len() == 0 {
			sb.WriteString("ไม่มีข้อมูล")
		}
		message.Messages[0].Text = sb.String()

	}

	replyMessageLine(message)

}

var richMenu = make(map[string]string)

func init() {
	initRichMenu()
}

func initRichMenu() {
	richMenu["home"] = "richmenu-f84842b69cb5bd69b414e7d3be367a91"

}
func replyMessageLine(Message ReplyMessage) error {
	value, _ := json.Marshal(Message)

	url := "https://api.line.me/v2/bot/message/reply"

	var jsonStr = []byte(value)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+"iGTkDJ78gP4im/HgeU7PG4ycQZM1scdSsEA413EWaY/lhpEwTuUeY73XD3K8SQkzDZunSM2h/F9kdewS/TDwaPpkZcfADM5HZW02PzC8KY4jUPuTHG2uKcIsDqrWQGtLBN+Xy+H3CZ1hOZdd7qGujAdB04t89/1O/w1cDnyilFU=")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	return err
}
