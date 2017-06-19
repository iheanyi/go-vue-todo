<template>
  <div class="todos">
    <h4 v-if="isFetching">Fetching Todos</h4>
    <div class="pa4" v-else>
        <input @keyup.enter="addTodo" v-model="description" type="text" placeholder="Add Todo" />
        <button @click="addTodo">+</button>
      <ul class="pt3">
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
      description: '',
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
    async addTodo() {
      const payload = {
        description: this.description,
      };

      const response = await axios.post('/api/todos', payload);
      const todo = response.data;
      const newTodos = [...this.todos, todo];

      this.todos = newTodos;
      this.description = '';
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
