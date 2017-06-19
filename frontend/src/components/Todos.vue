<template>
  <div class="todos">
    <h4 v-if="isFetching">Fetching Todos</h4>
    <div class="pa4" v-else>
        <input type="text" />
      <ul>
        <li class="todo-item" v-for="todo in todos">
          {{ todo.description }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'todos',
  data() {
    return {
      todos: [],
      isFetching: true,
    };
  },
  async created() {
    this.fetchTodos();
  },
  methods: {
    async fetchTodos() {
      const response = await axios.get('/api/todos');
      this.todos = response.data.todos;
      this.isFetching = false;
    },
  },
};
</script>

<style scoped>
.todo-item {
  list-style: none;
}

.todos {
  border-radius: 4px;
  border: 1px #eee solid;
  width: 50%;
  margin: 0 auto;
}
</style>
