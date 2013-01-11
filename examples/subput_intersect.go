package main

import (
  "fmt"
  "github.com/zond/god/client"
  "github.com/zond/god/setop"
)

func main() {
  conn := client.MustConn("localhost:9191")
  followersKey := []byte("mail@domain.tld/followers")
  followeesKey := []byte("mail@domain.tld/followees")
  conn.SubPut(followersKey, []byte("user1@domain.tld"), nil)
  conn.SubPut(followersKey, []byte("user2@domain.tld"), nil)
  conn.SubPut(followersKey, []byte("user3@domain.tld"), nil)
  conn.SubPut(followeesKey, []byte("user3@domain.tld"), nil)
  conn.SubPut(followeesKey, []byte("user4@domain.tld"), nil)
  for _, friend := range conn.SetExpression(setop.SetExpression{
    Code: fmt.Sprintf("(I %v %v)", string(followersKey), string(followeesKey)),
  }) {
    fmt.Println(string(friend.Key))
  }
}

// output: user3@domain.tld
