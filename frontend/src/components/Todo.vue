<template>
  <div>
    <div class="bb ph3 pv4 b--light-gray todo-item">
      <span class="b f6">{{ todo.createdAt | datetime }}</span>
      <div>
        {{ todo.description }}
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    props: ['todo'],
    data() {
      return {};
    },
  };
</script>

<style scoped>

</style>
