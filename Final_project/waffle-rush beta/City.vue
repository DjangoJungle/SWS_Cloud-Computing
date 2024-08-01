<template>
  <div>
    <div class="main" style="overflow-y: hidden; ">
      <el-header class="title">
        <div style="margin-top: 12px; display: inline-block;">
          <svg color="white" xmlns="http://www.w3.org/2000/svg" width="40" height="40" fill="currentColor" class="qi-102" viewBox="0 0 16 16">
            <path d="M13.268 15.162a.224.224 0 0 0-.233.042A2.99 2.99 0 0 1 11 16a2.99 2.99 0 0 1-2.035-.796.224.224 0 0 0-.233-.042 2 2 0 1 1-.383-3.832.224.224 0 0 0 .22-.087A2.996 2.996 0 0 1 11 10c1 0 1.887.49 2.432 1.243.05.069.136.102.22.087a2 2 0 1 1-.383 3.832ZM11 15.25c.752 0 1.418-.37 1.827-.936.086-.12.273-.134.388-.041a1.25 1.25 0 1 0 .209-2.082c-.132.068-.312.017-.373-.118a2.25 2.25 0 0 0-4.102 0c-.06.135-.241.186-.372.118a1.25 1.25 0 1 0 .209 2.082c.114-.093.301-.079.387.04A2.247 2.247 0 0 0 11 15.25ZM7.655 2.357a.5.5 0 0 0 .854-.353v-1.5a.5.5 0 1 0-1 0v1.5a.5.5 0 0 0 .146.353Zm-4.08 1.861c.06.026.126.039.191.039l.001-.001a.5.5 0 0 0 .355-.855l-1.064-1.06a.5.5 0 0 0-.707.708l1.062 1.06a.498.498 0 0 0 .162.11ZM.503 8.496h1.5a.5.5 0 1 0 0-1h-1.5a.5.5 0 0 0 0 1Zm1.914 5.221a.501.501 0 0 0 .631-.063l1.063-1.06a.5.5 0 0 0-.708-.707l-1.062 1.06a.5.5 0 0 0 .076.77ZM12.393 9a4.5 4.5 0 1 0-7.033 2.64l.718-.718A3.501 3.501 0 0 1 4.505 8a3.504 3.504 0 0 1 3.5-3.5A3.5 3.5 0 0 1 11.359 9h1.034Zm1.609-.49h1.5a.5.5 0 1 0 0-1h-1.5a.5.5 0 1 0 0 1Zm-2.031-4.327a.5.5 0 0 0 .633-.063l1.06-1.06a.5.5 0 1 0-.708-.708l-1.06 1.06a.5.5 0 0 0 .075.77Z"/>
          </svg>
          <span style="font-size: large; font-family: 'Microsoft YaHei'; color: #ffffff; font-weight: bold;">Cloud on Cloud</span>
        </div >
        <div style="margin-top: 12px; display: inline-block;">
          <router-link to="/MainMap">
            <el-button type="info">
              Map
            </el-button>
          </router-link>
        </div >
      </el-header>

      <div>
        <el-card style="width: 45%" title="city weather"  >
          <p style="padding: 2.5px;"><span style="font-weight: bold;">province：</span>{{ cityList.province }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">city：</span>{{ cityList.city }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">reportTime：</span>{{ cityList.reportTime }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">weather：</span>{{ cityList.weather }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">humidity：</span>{{ cityList.humidity }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">humidity_float：</span>{{ cityList.humidity_float }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">wind_direction：</span>{{ cityList.wind_direction }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">wind_power：</span>{{ cityList.wind_power }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">temperature：</span>{{ cityList.temperature }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">temperature_float：</span>{{ cityList.temperature_float }}</p>
        </el-card>
        <el-card style="width: 45%; margin-top: 10px; " title="predict data"  >
          <p style="padding: 2.5px;"><span style="font-size: large;  font-weight: bold;">Predict weather information in one hour:</span></p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">weather：</span>{{ predictList.weather }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">humidity_float：</span>{{ predictList.humidity_float }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">wind_direction：</span>{{ predictList.wind_direction }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">wind_power：</span>{{ predictList.wind_power }}</p>
          <p style="padding: 2.5px;"><span style="font-weight: bold;">temperature_float：</span>{{ predictList.temperature_float }}</p>
        </el-card>
        <el-card style="width: 45%; margin-top: 10px;">
          <p style="padding: 2.5px;" v-if="currentSuggestion === 'wear'"><span style="font-weight: bold;">What to wear today? ：</span>{{ suggestionList.wear }} will be good.</p>
          <p style="padding: 2.5px;" v-if="currentSuggestion === 'go_out'"><span style="font-weight: bold;">Is it good to go out? ：</span>{{ suggestionList.go_out }}</p>
          <p style="padding: 2.5px;" v-if="currentSuggestion === 'car'"><span style="font-weight: bold;">How about wash a car? ：</span>{{ suggestionList.car }}</p>
          <p style="padding: 2.5px;" v-if="currentSuggestion === 'umbrella'"><span style="font-weight: bold;">Do I need to take an umbrella with me? ：</span>{{ suggestionList.umbrella }}</p>
          <el-button style="margin-right: 0;" type="text" @click="showNextSuggestion">NEXT</el-button>
        </el-card>
      </div>
    </div>
  </div>

    <div class="chatroom" @mousedown="startDrag" :style="{ top: top + 'px', left: left + 'px' }">
      <div class="header">
        Chat Room: {{ currentRoom }}
        <span class="online-users">Online Users: {{ onlineUsers }}</span>
      </div>
      <div class="highlighted-toggle" @click="toggleHighlighted">
        Starred Messages
      </div>
      <div class="highlighted-messages" v-if="showHighlightedMessages">
        <div v-for="(message, index) in highlightedMessages" :key="index" class="message highlighted">
          {{ message.content }} ({{ message.likes }} likes) - {{ formatTimestamp(message.timestamp) }}
        </div>
      </div>
      <div class="messages">
        <div v-for="(message, index) in sortedMessages" :key="index" :class="{'word-my': message.userId === userId, 'word': message.userId !== userId}">
          <div v-if="message.userId !== userId" class="word">
            <img :src="getAvatarUrl(message.userId)" alt="Avatar">
            <div class="info">
              <p class="time">{{ message.userId }}  {{ formatTimestamp(message.timestamp) }}</p>
              <div class="info-content">{{ message.content }}
                <button @click="toggleLike(message)">👍 {{ message.likes }}</button>
              </div>
            </div>
          </div>
          <div v-else class="word-my">
            <div class="info">
              <p class="time">{{ message.userId }}  {{ formatTimestamp(message.timestamp) }}</p>
              <div class="info-content">
                <span v-for="(part, index) in formatContent(message.content)" :key="index">
                  <template v-if="isUrl(part)">
                    <a :href="part" target="_blank">{{ part }}</a>
                  </template>
                  <template v-else>
                    {{ part }}
                  </template>
                </span>
                <button @click="toggleLike(message)">👍 {{ message.likes }}</button>
              </div>
            </div>
            <img :src="getAvatarUrl(message.userId)" alt="Avatar">
          </div>
        </div>
      </div>
      <input v-model="newMessage" @keyup.enter="sendMessage" placeholder="Type your message here..." />
      <div v-if="errorMessage" class="error-message">{{ errorMessage }}</div>
    </div>
    <div class="sidebar">
      <input v-model="search" @input="filterRooms" placeholder="Search chatrooms" />
      <ul>
        <li v-for="room in filteredRooms" :key="room" @click="switchRoom(room)">
          {{ room }}
        </li>
      </ul>
    </div>

    <el-dialog v-model="showAlert">
      <strong>Warning：</strong> There's a climate warning in your city!
      <p style="padding: 2.5px;"><span style="font-weight: bold;">report time：</span>{{ warningList[0].reportTime }}</p>
      <p style="padding: 2.5px;"><span style="font-weight: bold;">type name：</span>{{ warningList[0].typeName }}</p>
      <p style="padding: 2.5px;"><span style="font-weight: bold;">severity：</span>{{ warningList[0].severity }}</p>
      <p style="padding: 2.5px;"><span style="font-weight: bold;">content：</span>{{ warningList[0].text }}</p>
      <template #footer>
                <span class="dialog-footer">
                    <el-button @click="showAlert = false">OK</el-button>
                </span>
      </template>
    </el-dialog>
</template>

<script>
import {Delete, Edit, Plus, Search} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'

//chatroom dependency
import { v4 as uuidv4 } from 'uuid';

export default {
  components: {Plus},
    data() {
        return {
          //chat room data
          socket: null,
          messages: [],
          highlightedMessages: [],
          newMessage: '',
          userId: '',
          likedMessages: [],
          errorMessage: '',
          top: 45,
          left: 650,
          dragging: false,
          offsetX: 0,
          offsetY: 0,
          showHighlightedMessages: false,
          currentRoom: 'Beijing',
          search: '',
          chatrooms: ['Beijing', 'Shanghai', 'Guangzhou', 'Shenzhen', 'Hangzhou'],
          filteredRooms: ['Beijing', 'Shanghai', 'Guangzhou', 'Shenzhen', 'Hangzhou'],
          onlineUsers: 0,
              // 增加消息内容解码功能
          decodeContent: decodeURIComponent,
          encodeContent: encodeURIComponent,

          //weather data
          showAlert: false,
          get_province : '',
          get_city : '',
          cityList:{
            province : '',
            city : '',
            reportTime : '',
            weather : '',
            temperature : 0,
            temperature_float : 0.0,
            humidity : 0,
            humidity_float : 0.0,
            wind_direction : '',
            wind_power : 0,
          },
          predictList:{
            province : '',
            city : '',
            reportTime : '',
            weather : '',
            temperature_float : 0.0,
            humidity_float : 0.0,
            wind_direction : '',
            wind_power : 0,
          },
          suggestionList:{
            province : '',
            city : '',
            reportTime : '',
            wear : '',
            umbrella : '',
            car : '',
            go_out : '',
          },
          warningList:[{
            province : '',
            city : '',
            reportTime : '',
            startTime : '',
            endTime : '',
            title : '',
            status : '',
            level : '',
            severity : '',
            severityColor : '',
            typeName:'',
            text:'',
          }],
          currentSuggestion: 'wear', // 初始显示穿衣建议
          suggestionOrder: ['wear', 'go_out', 'car', 'umbrella'], // 生活指南项顺序
        }
    },

  //chat room functions
  computed: {
    sortedMessages() {
      return [...this.messages].sort((a, b) => a.timestamp - b.timestamp);
    }
  },
  created() {
    this.socket = new WebSocket('/ws'); // 连接到Java WebSocket服务器

    // 获取或生成用户ID
    this.userId = localStorage.getItem('userId') || uuidv4();
    localStorage.setItem('userId', this.userId);

    // 获取已经点赞的消息
    this.likedMessages = JSON.parse(localStorage.getItem('likedMessages')) || [];

    this.socket.onmessage = (event) => {
      console.log('Received message:', event.data);
      const parts = event.data.split(':');
      if (parts[0] === 'NEW') {
        this.messages.push({
          id: parts[1],
          userId: parts[2],
          content: decodeURIComponent(parts[3]), // 使用解码
          likes: parseInt(parts[4], 10), // 确保点赞数是数字
          timestamp: parseInt(parts[5], 10) // 确保时间戳是数字
        });
      } else if (parts[0] === 'UPDATE') {
        const msg = this.messages.find(m => m.id === parts[1]);
        if (msg) {
          msg.likes = parseInt(parts[2], 10);
          if (msg.likes >= 5 && !this.highlightedMessages.some(m => m.id === msg.id)) {
            this.highlightedMessages.push(msg);
          } else if (msg.likes < 5) {
            this.highlightedMessages = this.highlightedMessages.filter(m => m.id !== msg.id);
          }
        }
      } else if (parts[0] === 'HIGHLIGHT') {
        if (!this.highlightedMessages.some(m => m.id === parts[1])) {
          this.highlightedMessages.push({
            id: parts[1],
            content: decodeURIComponent(parts[2]), // 使用解码
            likes: parseInt(parts[3], 10),
            timestamp: parseInt(parts[4], 10)
          });
        }
      } else if (parts[0] === 'ONLINE_USERS') {
        this.onlineUsers = parseInt(parts[1], 10);
      } else if (parts[0] === 'ERROR') {
        this.errorMessage = parts.slice(1).join(':');
        setTimeout(() => {
          this.errorMessage = '';
        }, 5000);
      }
    };

    this.socket.onopen = () => {
      console.log('WebSocket connection opened');
      this.switchRoom(this.currentRoom); // 加载初始聊天室的消息
    };

    this.socket.onclose = () => {
      console.log('WebSocket connection closed');
    };

    this.socket.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    window.addEventListener('mousemove', this.onDrag);
    window.addEventListener('mouseup', this.stopDrag);
  },
  unmounted() {
    window.removeEventListener('mousemove', this.onDrag);
    window.removeEventListener('mouseup', this.stopDrag);
  },
    methods: {
      QueryCityInfo() {
        this.get_province = sessionStorage.getItem('province')
        this.get_city = sessionStorage.getItem('city')

        axios.get('/gci',{
          params: {
            province: this.get_province,
            city: this.get_city
          }})
            .then(response => {
              if (response.status === 200) {
                console.log("success");
                this.cityList = response.data.payload;
                console.log(this.cityList);
              }
            })
            .catch(error => {
              console.log("数据获取错误，请联系开发人员");
            })
      },
      QueryPredictInfo() {
        this.get_province = sessionStorage.getItem('province')
        this.get_city = sessionStorage.getItem('city')

        axios.get('/gcpi',{
          params: {
            province: this.get_province,
            city: this.get_city
          }})
            .then(response => {
              console.log("success");
              if (response.status === 200) {
                console.log("success and 200 sc");
                this.predictList = response.data.payload;
                console.log(this.predictList);
              }
            })
            .catch(error => {
              console.log("数据获取错误，请联系开发人员");
            })
      },

      /********************************************新增请求******************************************/
      QueryWarningInfo() {
        this.get_province = sessionStorage.getItem('province')
        this.get_city = sessionStorage.getItem('city')
        axios.get('/gcw',{
          params: {
            province: this.get_province,
            city: this.get_city
          }})
            .then(response => {
              if (response.status === 200) {
                this.warningList = response.data.payload;
                console.log(this.warningList);
                this.showAlert = true;
              }else if(response.status === 404){
                console.log("No warning in this city.");
              }
            })
            .catch(error => {
              console.log("数据获取错误，请联系开发人员");
            })
      },

      /********************************************新增请求******************************************/
      QuerySuggestionInfo() {
        this.get_province = sessionStorage.getItem('province')
        this.get_city = sessionStorage.getItem('city')
        axios.get('gcs',{
          params: {
            province: this.get_province,
            city: this.get_city
          }})
            .then(response => {
              if (response.status === 200) {
                this.suggestionList = response.data.payload;
                console.log(this.suggestionList);
              }
            })
            .catch(error => {
              console.log("数据获取错误，请联系开发人员");
            })
      },
      showNextSuggestion() {
        const currentIndex = this.suggestionOrder.indexOf(this.currentSuggestion);
        if (currentIndex === -1 || currentIndex === this.suggestionOrder.length - 1) {
          this.currentSuggestion = this.suggestionOrder[0]; // 如果当前是最后一项，切换到第一项
        } else {
          this.currentSuggestion = this.suggestionOrder[currentIndex + 1]; // 切换到下一项
        }
      },

      //chat room methods
      sendMessage() {
        if (this.newMessage.trim() !== '') {
          const messageId = uuidv4();
          const encodedMessage = encodeURIComponent(this.newMessage); // 使用编码
          this.socket.send(messageId + ':' + this.userId + ':' + this.currentRoom + ':' + encodedMessage);
          this.newMessage = '';
        }
      },
      toggleLike(message) {
        if (this.likedMessages.includes(message.id)) {
          this.socket.send('UNLIKE:' + message.id);
          this.likedMessages = this.likedMessages.filter(id => id !== message.id);
        } else {
          this.socket.send('LIKE:' + message.id);
          this.likedMessages.push(message.id);
        }
        localStorage.setItem('likedMessages', JSON.stringify(this.likedMessages));
      },
      formatTimestamp(timestamp) {
        const date = new Date(Number(timestamp));
        return date.toLocaleString();
      },
      formatContent(content) {
        const urlPattern = /(https?:\/\/[^\s]+)/g;
        return content.split(urlPattern);
      },
      isUrl(text) {
        const urlPattern = /^https?:\/\/[^\s]+$/;
        return urlPattern.test(text);
      },
      getAvatarUrl(userId) {
        const avatars = ['avatar1.png', 'avatar2.png', 'avatar3.png', 'avatar4.png', 'avatar5.png'];
        const index = parseInt(userId, 16) % avatars.length;
        const avatarUrl = `/avatars/${avatars[index]}`;
        console.log(`Generated avatar URL for userId ${userId}: ${avatarUrl}`);
        return avatarUrl;
      },
      startDrag(event) {
        if (event.target.className === 'header') { // 仅在点击header时拖动
          this.dragging = true;
          this.offsetX = event.clientX - this.left;
          this.offsetY = event.clientY - this.top;
        }
      },
      onDrag(event) {
        if (this.dragging) {
          this.top = event.clientY - this.offsetY;
          this.left = event.clientX - this.offsetX;
        }
      },
      stopDrag() {
        this.dragging = false;
      },
      toggleHighlighted() {
        this.showHighlightedMessages = !this.showHighlightedMessages;
      },
      filterRooms() {
        this.filteredRooms = this.chatrooms.filter(room => room.includes(this.search));
      },
      switchRoom(room) {
        this.currentRoom = room;
        this.messages = [];
        this.highlightedMessages = [];
        this.socket.send('CHATROOM:' + room);
      },
    },
    mounted() { 
        this.QueryCityInfo();
        this.QueryPredictInfo();
        this.QuerySuggestionInfo();
        this.QueryWarningInfo();
    },
  
}

</script>


<style scoped>
.main {
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  min-height: 100%;
  height: auto;
  background-color: #dcdcdc;

}

.title {
  background-color: rgb(30, 31, 34);
  height: 60px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
}



.sidebar {
  right: 0;
  top: 60px;
  width: 200px;
  border-left: 1px solid #ccc; /* 将边框从右侧移到左侧 */
  padding: 10px;
  box-sizing: border-box;
  position: absolute;
  background-color: white;
}

.sidebar input {
  width: 100%;
  padding: 5px;
  margin-bottom: 10px;
  box-sizing: border-box;
}

.sidebar ul {
  list-style: none;
  padding: 0;
}

.sidebar li {
  padding: 5px 0;
  cursor: pointer;
}

.sidebar li:hover {
  background-color: #f0f0f0;
}

.chatroom {
  display: flex;
  flex-direction: column;
  border: 1px solid #ccc;
  background: white;
  position: relative;
  width: 600px; /* 固定聊天窗口宽度 */
  height: 650px; /* 固定聊天窗口高度 */
}

.header {
  background: #f1f1f1;
  padding: 10px;
  cursor: move;
  text-align: center;
  font-weight: bold;
  border-bottom: 1px solid #ccc;
}

.online-users {
  float: right;
}

.highlighted-toggle {
  background: #f1f1f1;
  padding: 10px;
  cursor: pointer;
  text-align: center;
  font-weight: bold;
  border-bottom: 1px solid #ccc;
  margin-bottom: 10px;
}

.highlighted-messages {
  background-color: #f9f9f9;
  padding: 10px;
}

.messages {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
}

.word, .word-my {
  display: flex;
  margin-bottom: 20px;
}

.word img, .word-my img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

.word .info, .word-my .info {
  margin-left: 10px;
}

.word .info .time, .word-my .info .time {
  font-size: 12px;
  color: rgba(51, 51, 51, 0.8);
  margin: 0;
  height: 20px;
  line-height: 20px;
  margin-top: -5px;
}

.word .info .info-content, .word-my .info .info-content {
  padding: 10px;
  font-size: 14px;
  position: relative;
  margin-top: 8px;
  background: #fff;
  word-wrap: break-word; /* 自动换行 */
  white-space: pre-wrap; /* 保留空格和换行符 */
  word-break: break-all; /* 长单词换行 */
}

.word .info .info-content::before {
  position: absolute;
  left: -8px;
  top: 8px;
  content: '';
  border-right: 10px solid #FFF;
  border-top: 8px solid transparent;
  border-bottom: 8px solid transparent;
}

.word-my {
  justify-content: flex-end;
}

.word-my .info {
  width: 90%;
  text-align: right;
}

.word-my .info .time {
  margin-right: 10px;
}

.word-my .info .info-content {
  max-width: 70%;
  float: right;
  margin-right: 10px;
  background: #A3C3F6;
  text-align: left;
}

.word-my .info .info-content::after {
  position: absolute;
  right: -8px;
  top: 8px;
  content: '';
  border-left: 10px solid #A3C3F6;
  border-top: 8px solid transparent;
  border-bottom: 8px solid transparent;
}

input {
  padding: 10px;
  border: none;
  border-top: 1px solid #ddd;
  width: 100%;
  box-sizing: border-box;
}

button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
  margin-left: 10px;
  font-size: 1.2em;
}

.error-message {
  color: red;
  font-weight: bold;
  margin-top: 10px;
}
</style>