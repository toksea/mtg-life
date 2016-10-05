# MTG 对战平台（类似 VPT）

## 技术

* 前端
  * webpack
  * react
  * redux
  * websocket（浏览器原生）
  * 基于 https://github.com/reactjs/redux/tree/master/examples/real-world
* 后端
  * golang
  * https://github.com/gorilla/websocket
  * 基于 https://github.com/gorilla/websocket/tree/master/examples/chat

## 数据结构

* match 轮
* game 局（三局两胜）
* player
* deck
* zone
    * battleground
    * graveyard
    * exiled zone
    * commander zone
