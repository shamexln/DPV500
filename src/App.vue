<template>
  <div id="app" ref="document">
     <MyHeader :addTodo="addTodo"></MyHeader>
        <router-view></router-view>
        <div>
          <label>参数记录</label>
        </div>
        <img class="profile-image" :src="userImage" />

        <div v-if="!userImage">
          <input
            type="file"
            round
            class="change-profile-image"
            @change="onFileChange"
          />
        </div>
        <div v-else>
          <button
            class="delete-profile-image"
            color="secondary"
            icon="delete"
            @click="removeImage"
          >
            Delete
          </button>
        </div>
        <button @click="capture">capture</button>
        <button @click="show">show</button>
        <button @click="download1">Download to PDF</button>
        <button @click="exportToPDF">Export to PDF</button>
        <div class="field" v-for="(item, index) in imgList" :key="index">
          <div>
          
            <img :src="require('@/assets/imgs' + item.slice(1))" alt="" />
       
          
          </div>
          <button @click="removeImg(index)">Delete</button>
        </div>
   <MyFooter :todos="todos"></MyFooter> 
  </div>
</template>

<script>
//import axios from 'axios'
import * as jsPDF from "jspdf";
import html2canvas from "html2canvas";
import MyHeader from "./components/MyHeader";
// import MyList from "./components/MyList";
import MyFooter from "./components/MyFooter";
import html2pdf from 'html2pdf.js'

import { captureNewImage, captureDeleteImage } from "@/api";
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
      userImage: "",
      todos: [
        { id: "001", title: "Play Game", done: true },
        { id: "002", title: "Play Ball", done: false },
        { id: "003", title: "Play Car", done: true },
      ],
    };
  },
  methods: {
    		exportToPDF () {
          let body = document.body
      let html = document.documentElement
      let height = Math.max(body.scrollHeight, body.offsetHeight,
                       html.clientHeight, html.scrollHeight, html.offsetHeight)
      let heightCM = height / 35.35

				html2pdf(this.$refs.document, {
					margin: 1,
					filename: 'document.pdf',
					image: { type: 'jpeg', quality: 0.98 },
					//html2canvas: { dpi: 300, letterRendering: true },
          html2canvas:  {scale: 1},
          pagebreak: {before: '.newPage', avoid: ['h1', 'h3', 'h4', 'textarea','.field']},
					jsPDF: { unit: 'cm', format: [heightCM, 60], orientation: 'landscape' }
				})
			},
    download1() {
      let windowHeight = window.innerHeight;
      let windowWidth = window.innerWidth;

      let pdf = new jsPDF.default();

      let canvasElement = document.createElement("canvas");
      canvasElement.width = windowWidth;
      canvasElement.height = windowHeight;

      html2canvas(this.$refs.document, {
        canvas: canvasElement,
        width: windowWidth,
        height: windowHeight,
      })
        .then((canvas) => {
          const img = canvas.toDataURL("image/jpeg", 1);
          //document.body.appendChild(canvas);
          pdf.addImage(img, "JPEG", 10, 10, 200, 250);
          pdf.save("sample.pdf");

          alert("works");
        })
        .catch((err) => {
          alert(err);
        });
    },
    //jspdf does not include the bootstrap style layout
    generatePdf() {
      var doc = new jsPDF.default();

      //  var doc = new jsPDF('p', 'pt', 'A4');
      //    let margins = {
      //       top: 80,
      //       bottom: 60,
      //       left: 40,
      //       width: 522
      //   };

      doc.html(this.$refs.testHtml, {
        html2canvas: {
          // insert html2canvas options here, e.g.
          width: 200,
        },
        callback: function () {
          window.open(doc.output("bloburl"));
        },
      });

      doc.save("test.pdf");
    },
    download() {
      const doc = new jsPDF();
      const contentHtml = this.$refs.content.innerHTML;
      doc.fromHTML(contentHtml, 15, 15, {
        width: 170,
      });
      doc.save("sample.pdf");
    },

    downloadWithCSS() {
      const doc = new jsPDF();
      /** WITH CSS */
      var canvasElement = document.createElement("canvas");
      html2canvas(this.$refs.content, { canvas: canvasElement }).then(function (
        canvas
      ) {
        const img = canvas.toDataURL("image/jpeg", 0.8);
        doc.addImage(img, "JPEG", 20, 20);
        doc.save("sample.pdf");
      });
    },
    removeImg: function (index) {
      let fileName = this.imgList[index];
      let newfileName = fileName.substring(0, fileName.lastIndexOf("."));
      let resdata = captureDeleteImage(newfileName);
      console.log(resdata);
    },
    show() {
      const files = require
        .context("@/assets/imgs", true, /\.*\.jpg|jpeg|png$/)
        .keys();
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
    generatePDF() {},

    onFileChange(e) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length) {
        return;
      }
      this.createImage(files[0]);
    },
    createImage(file) {
      var reader = new FileReader();
      var vm = this;

      reader.onload = (e) => {
        vm.userImage = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    // eslint-disable-next-line no-unused-vars
    removeImage: function (e) {
      this.userImage = "";
    },
    addTodo(todoObj) {
      this.todos.unshift(todoObj);
    },
    deleteTodo(id) {
      this.todos = this.todos.filter((item) => item.id !== id);
    },
    checkTodo(id) {
      this.todos.forEach((item) => {
        if (item.id === id) {
          item.done = !item.done;
        }
      });
    },
  },
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
.field {

}

</style>
