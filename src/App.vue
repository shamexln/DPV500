<template>
  <div id="app">
    <MyHeader :addTodo=addTodo></MyHeader>
    <MyList :todos=todos :checkTodo=checkTodo :deleteTodo=deleteTodo></MyList> -->
     <router-view></router-view>
    <MyFooter :todos=todos></MyFooter>
    
    <br/>
    <br/>
    <br/>
    <div>
      <label>参数记录</label>
    </div>
    <img class="profile-image" :src="userImage" />

    <div v-if="!userImage">
      <input type="file" round class="change-profile-image" @change="onFileChange" />
    </div>
    <div v-else>
      <button class="delete-profile-image" color="secondary" icon="delete" @click="removeImage">Delete</button>
    </div>
    <button @click="capture">capture</button>
    <button @click="show">show</button>
    <div v-for="(item,index) in imgList" :key="index">
      <div>
        <img :src="require('@/assets/imgs'+item.slice(1))" alt=""  />
        
      </div>
      <button @click="removeImg(index)">Delete</button>
    </div>
    
  </div>
   
</template>

<script>

// import jspdf from 'jspdf'
//import axios from 'axios'
import MyHeader from "./components/MyHeader";
// import MyList from "./components/MyList";
import MyFooter from "./components/MyFooter";
import {captureNewImage, captureDeleteImage} from '@/api';
export default {
  name: "App",
  components: {
    MyHeader,
    // MyList,
     MyFooter,
  },
  data() {
    return {
      imgList: [],
      userImage: '',
      todos: [
        { id: "001", title: "Play Game", done: true },
        { id: "002", title: "Play Ball", done: false },
        { id: "003", title: "Play Car", done: true },
      ],
    };
  },
  methods:{
    removeImg: function (index) {
      let fileName = this.imgList[index]
      let newfileName = fileName.substring(0,fileName.lastIndexOf('.'))
      let resdata = captureDeleteImage(newfileName);
      console.log(resdata);
    },
    show () {
      const files = require.context("@/assets/imgs", true, /\.*\.jpg|jpeg|png$/).keys();
      this.imgList = files;
    },
    capture() {
      // get files from backend by axos
      // create new image
      // axios
      // .post('images')
      // .then((response) => {
      //   console.log(response.data)
      // })
      let resdata = captureNewImage({});
      console.log(resdata);
      
    },
    generatePDF() {

    },

    onFileChange(e) {
        var files = e.target.files || e.dataTransfer.files
        if (!files.length) {
          return
        }
        this.createImage(files[0])
      },
      createImage(file) {
        var reader = new FileReader()
        var vm = this

        reader.onload = (e) => {
          vm.userImage = e.target.result
        }
        reader.readAsDataURL(file)
      },
       // eslint-disable-next-line no-unused-vars
       removeImage: function (e) {
         this.userImage = ''
       },
    addTodo(todoObj) {
      this.todos.unshift(todoObj)
    },
    deleteTodo(id) {
      this.todos = this.todos.filter( item => item.id !== id )
    },
    checkTodo(id) {
      this.todos.forEach((item) => {
        if (item.id === id) {
          item.done = !item.done
        }
      })
    }
  }
};
</script>

<style>
body {
  background: #fff;
}

.btn {
  display: inline-block;
  padding: 4px 12px;
  margin-bottom: 0;
  font-size: 14px;
  line-height: 20px;
  text-align: center;
  vertical-align: middle;
  cursor: pointer;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2),
    0 1px 2px rgba(0, 0, 0, 0.05);
  border-radius: 4px;
}

.btn-danger {
  color: #fff;
  background-color: #da4f49;
  border: 1px solid #bd362f;
}

.btn-danger:hover {
  color: #fff;
  background-color: #bd362f;
}

.btn:focus {
  outline: none;
}

.todo-container {
  width: 600px;
  margin: 0 auto;
}
.todo-container .todo-wrap {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
}
</style>
