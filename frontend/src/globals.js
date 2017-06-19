import Vue from 'vue';
import moment from 'moment';

Vue.filter('datetime', (value) => {
  const formattedDate = moment(value).format('MMMM Do YYYY, h:mm:ss a');
  return formattedDate;
});
