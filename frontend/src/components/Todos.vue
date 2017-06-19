<template>
  <div class="todos">
    <h4 class="pa4" v-if="isFetching">Fetching Todos</h4>
    <div v-else>
      <div class="flex">
        <input class="bn bb bg-near-white bw1 w-90 black bold h3 pa3" @keyup.enter="addTodo" v-model="description" type="text" placeholder="Add Todo" />
        <button class="w-10 bg-navy white bn" @click="addTodo">+</button>
      </div>
      <ul class="pa3">
        <li class="pv2 bv todo-item" v-for="todo in todos">
          <div>
            {{ todo.description }}
          </div>
          <span>{{ todo.createdAt | datetime }}</span>
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
