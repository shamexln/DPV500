<template>
  <div class="field">
  <label>报告总结</label>
  <br/>

  <textarea cols="150" rows="10" v-model="msg"></textarea>
      
  </div>
</template>

<script>
export default {
    name:'MyFooter',
    props:['todos'],
    computed: {
        total() {
            return this.todos.length
        },
        doneTotal() {
            return this.todos.reduce((pre, currentVal) => {
                return pre + (currentVal.done ? 1 : 0)
            }, 0)
            
        }
    }

}
</script>

<style>
.field {

}
</style>