<template>
  <div class="bb pa3 b--light-gray todo-item">
    <span class="db mb2 b f6">{{ editTodo.createdAt | datetime }}</span>
    <span class="db gray f6">Updated: {{ editTodo.updatedAt | datetime }}</span>
    <div v-if="isEditing">
      <input v-focus type="text" v-model="editTodo.description" class="outline-transparent ba mv2 b--light-gray bw1 w-100 black bold pa2" />

      <div class="mb2">
        <button @click="onSaveClick" class="pointer bg-dark-blue bn dim dib white pv2 ph3 mr2">Save</button>
        <button @click="onCancelClick" class="pointer bg-dark-blue bn dim dib white pv2 ph3 mr2">Cancel</button>
      </div>
    </div>
    <div v-else class="pv2">
      {{ editTodo.description }}
    </div>
    <div class="f6">
      <button @click="onRemoveClick(todo.id)" class="pointer pa0 bg-transparent bn blue underline dib mr2">Delete</button>
      <button v-if="!isEditing" @click="onEditClick" class="pointer pa0 bg-transparent bn blue underline dib mr2">Edit</button>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'todo',
    props: ['todo', 'onRemoveClick'],
    data() {
      return {
        isEditing: false,
        editTodo: {
          ...this.todo,
        },
        cachedTodo: null,
      };
    },
    methods: {
      onEditClick() {
        this.isEditing = true;
        this.cachedTodo = {
          ...this.editTodo,
        };
      },
      onCancelClick() {
        this.isEditing = false;
        this.editTodo = {
          ...this.cachedTodo,
        };
      },
      async onSaveClick() {
        const response = await axios.put(`/api/todos/${this.todo.id}`, this.editTodo);
        this.editTodo = response.data;
        this.isEditing = false;
      },
    },
  };
</script>

<style scoped>
</style>
