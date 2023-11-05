package main

import (
	"chat-box/app/routes"

	"github.com/gin-gonic/gin"
)

type MyData struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	UserId  string `json:"user_id"`
	Connect bool   `json:"connect"`
}

func addToSliceIfNotExists(slice map[string]string, key string, str string) map[string]string {
	for _, s := range slice {
		if s == str {
			return slice // Chuỗi đã tồn tại trong slice, không thêm nữa
		}
	}
	slice[key] = str
	return slice // Chuỗi chưa tồn tại, thêm vào slice
}
func main() {
	r := gin.Default()
	// content, err := ioutil.ReadFile("./app/data/user.txt")
	// if err != nil {
	// 	fmt.Println("Không thể đọc tệp:", err)
	// 	return
	// }
	// fmt.Println(content)
	// m := melody.New()
	// listUser := make([]string, 1)
	// mySlice := map[string]string{}
	// messages := map[string][]string{}
	routes.Web(r)
	// 	c.HTML(http.StatusOK, "index.html", data)
	// })

	// r.GET("/ws", func(c *gin.Context) {
	// 	m.HandleRequest(c.Writer, c.Request)
	// })
	// r.GET("/ws/:user/:id", func(c *gin.Context) {
	// 	data := map[string]interface{}{
	// 		"id": c.Param("id"),
	// 	}

	// 	m.HandleRequestWithKeys(c.Writer, c.Request, data)
	// })
	// m.HandleDisconnect(func(s *melody.Session) {
	// 	fmt.Println("disconnect")
	// })
	// m.HandleConnect(func(s *melody.Session) {
	// 	fmt.Println("connect")
	// 	id, ok := s.Keys["id"].(string)
	// 	if !ok {
	// 		return
	// 	}
	// 	s.Set("group"+s.RemoteAddr().String(), id)
	// 	s.Set("user"+s.RemoteAddr().String(), id)
	// 	var listSessionss []*melody.Session
	// 	listSessionss = append(listSessionss, s)
	// 	// send message cho người mới
	// 	if data, ok := messages[id]; ok {
	// 		for _, ms := range data {
	// 			fmt.Println(ms)
	// 			m.BroadcastMultiple([]byte(ms), listSessionss)
	// 		}
	// 	}

	// 	//  send tất cả người dùng về có người mới
	// 	// sessions, _ := m.Sessions()
	// 	// var listSessions []*melody.Session

	// 	// m.BroadcastMultiple([]byte(jsonData), listSessions)
	// })

	// m.HandleMessage(func(s *melody.Session, msg []byte) {
	// 	// fmt.Println(msg)
	// 	str := string(msg)
	// 	// fmt.Println(str)
	// 	// jsonData := `{"÷d": 12, "message": "âsd"}`
	// 	var data MyData
	// 	err := json.Unmarshal([]byte(str), &data)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	// 	jsonData, _ := json.Marshal(data)
	// 	messages[data.UserId] = append(messages[data.UserId], string(jsonData))
	// 	sessions, _ := m.Sessions()
	// 	var listSessions []*melody.Session
	// 	for _, session := range sessions {
	// 		id, _ := session.Get("group" + session.RemoteAddr().String())
	// 		// f.t
	// 		if data.UserId == id {
	// 			listSessions = append(listSessions, session)
	// 		}
	// 	}
	// 	m.BroadcastMultiple([]byte(jsonData), listSessions)
	// })

	r.Run(":5000")
}
