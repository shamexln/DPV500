<template>
  <div>
    <div>
      <label class="word-v-middle">胸部电阻抗断层</label>
    </div>
    <div>
      <label>患者信息</label>
    </div>
    <div>
      <label>床号</label>
      <input type="text" v-model="newItem.chuanghao" />
      <label>姓名</label>
      <input type="text" v-model="newItem.xingming" />
      <label>性别</label>
      <input type="text" v-model="newItem.xingbie" />
      <label>年龄</label>
      <input type="text" v-model="newItem.nianling" />
      <label>住院号</label>
      <input type="text" v-model="newItem.zhuyuanhao" />
      <label>入院诊断</label>
      <input type="text" v-model="newItem.ruyuanzhenduan" />
      <label>入住ICU时间</label>
      <input type="text" v-model="newItem.ruyuanicu" />
      <label>检测时间</label>
      <input type="text" v-model="newItem.jianceshijian" />
      <label>患者体位</label>
      <input type="text" v-model="newItem.tiwei" />
      <label>体位时长</label>
      <input type="text" v-model="newItem.shichang" />
    </div>
  </div>
</template>

<script>
import { nanoid } from "nanoid";
export default {
  name: "MyHeader",
  data() {
    return {
      newItem: {
        chuanghao: "",
        xingming: "",
        xingbie: "",
        nianling: "",
        zhuyuanhao: "",
        ruyuanzhenduan: "",
        ruyuanicu: "",
        jianceshijian: "",
        tiwei: "",
        shichang: "",
      },
    };
  },
  methods: {
    add() {
      const todoObj = { id: nanoid(), title: this.newItem, done: false };
      this.addTodo(todoObj);
      this.newItem = "";
    },
  },
  props: ["addTodo"],
};
</script>

<style>
/*header*/
.todo-header input {
  width: 560px;
  height: 28px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 4px 7px;
}

.todo-header input:focus {
  outline: none;
  border-color: rgba(82, 168, 236, 0.8);
  box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075),
    0 0 8px rgba(82, 168, 236, 0.6);
}
.word-v-middle{

margin-bottom: 0;

font-size: 12px;

min-height: 31px;

display: flex;

align-items: center;

justify-content: center;

height: 31px;

margin-top: 5px;

color: #87878a;

white-space: normal;

}

.word-v-middle span{

text-align: left;

font-size: 10px;

text-overflow: -o-ellipsis-lastline;

overflow: hidden;

text-overflow: ellipsis;

display: -webkit-box;

-webkit-line-clamp: 2;

-webkit-box-orient: vertical;

}
</style>
