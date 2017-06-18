<template>
  <div class="todos">
    This is the Todos Compoent
  </div>
</template>

<script>

export default {
  name: 'todos',
  data() {
    return {
      todos: [
        {
          description: 'This is a test todo',
          isCompleted: false,

        },
      ],
    };
  },
};

</script>

<style scoped>

</style>
