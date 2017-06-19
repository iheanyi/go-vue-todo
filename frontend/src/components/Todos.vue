<template>
  <div class="w-100 w-50-ns centered todos">
    <h4 class="pa4" v-if="isFetching">Fetching Todos</h4>
    <div v-else>
      <div class="flex">
        <div class="bb b--light-gray bw1 w-100">
          <input class="bn b--light-gray bw1 w-100 black bold h3 pa3" @keyup.enter="addTodo" v-model="description" type="text" placeholder="Add Todo" />
        </div>
        <button class="w-10 bg-navy white bn" @click="addTodo">+</button>
      </div>
      <div v-if="todos.length">
        <todo
           v-for="todo in todos"
           :todo="todo"
           :onRemoveClick="deleteTodo"
           :key="todo.id"
        />
      </div>
      <div class="ph3 pv4" v-else>
        You haven't added any todos yet!
      </div>
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
      async deleteTodo(id) {
        await axios.delete(`/api/todos/${id}`);
        this.todos = this.todos.filter((todo) => {
          const todoId = todo.id;
          return todoId !== id;
        });
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
  margin: 0 auto;
}
</style>
