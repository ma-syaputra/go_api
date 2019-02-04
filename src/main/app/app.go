package app

import (
	//"fmt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"strconv"
)

func GetInstructions(c *gin.Context) {
	// c.JSON(200, gin.H{"ok": "GET api/v1/instructions"})
	data := 14
	response := fmt.Sprintf("Variable string %d content", data)
	fmt.Println(response)
}

func GetInstruction(c *gin.Context) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	client := redis.NewClient(&redis.Options{
		Addr:     "172.17.6.46:6379",
		Password: "",
		DB:       1,
	})
	first := c.Params.ByName("start")
	last := c.Params.ByName("stop")
	channel := c.Params.ByName("channel")
	channel_id, _ := strconv.ParseInt(channel, 0, 64)
	address := fmt.Sprintf("okz:breaking:%d", channel_id)
	start, _ := strconv.ParseInt(first, 0, 64)
	stop, _ := strconv.ParseInt(last, 0, 64)
	var define string = address
	val, err := client.LRange(define, start, stop-1).Result()
	if err != nil {
		c.JSON(200, gin.H{"message": "Data Empty"})
	}
	if len(val) > 0 {
		c.JSON(200, gin.H{"breaking": val})
	}
}

func PostInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "POST api/v1/instructions"})

}

func UpdateInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "PUT api/v1/instructions/1"})

}

func DeleteInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "DELETE api/v1/instructions/1"})
}
