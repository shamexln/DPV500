<template>
  <div>
      <input type="checkbox" v-model="isAll"/>
      <span>
			<span>已完成{{doneTotal}}</span> / 全部{{total}}
      </span>
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

</style>