import moment from 'moment';
import Vue from 'vue';
import Todo from './components/Todo';

Vue.component('todo', Todo);

Vue.filter('datetime', (value) => {
  const formattedDate = moment(value).format('MMMM Do YYYY, h:mm:ss a');
  return formattedDate;
});

Vue.directive('focus', {
  inserted(el) {
    el.focus();
  },
});
