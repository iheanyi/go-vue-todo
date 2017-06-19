<template>
  <div class="bb ph3 pv4 b--light-gray todo-item">
    <span class="b f6">{{ todo.createdAt | datetime }}</span>
    <div class="pv2">
      {{ todo.description }}
    </div>
    <div class="f6">
      <button @click="onRemoveClick(todo.id)" class="pointer pa0 bg-transparent bn blue underline dib mr2">Delete</button>
    </div>
  </div>
</template>

<script>
  export default {
    props: ['todo', 'onRemoveClick'],
    data() {
      return {};
    },
    methods: {},
  };
</script>

<style scoped>
</style>
