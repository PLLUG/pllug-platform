import Ember from 'ember';

export default Ember.Route.extend({
  actions: {
    save() {
      console.log("Data saved");
    }
  }
});
