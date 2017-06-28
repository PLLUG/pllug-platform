import Ember from 'ember';

const { Component, Object, computed } = Ember;

export default Ember.Component.extend({
  errors: Object.extend({
    email: [],
    password: []
  }),

  validData: computed.and('email', 'password'),

  actions: {
    save() {
      let { email, password } = this.getProperties('email', password);
      this.sendAction('onSubmit',  this.getProperties('email', 'password'));
    }
  }
});
