<style>
  .container {
    height: 100vh;
    width: 100vw;
    position: fixed;
    top: 0;
    left: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 500;

    background: rgba(0, 0, 0, .75);
  }

  .modal-container {
    width: 70vw;

    padding: 20px;

    border-radius: 15px;

    display: flex;
    flex-wrap: wrap;

    background: #fff;
  }

  .modal-row {
    width: 100%;
    margin-bottom: 20px;
  }

  .modal-half-row {
    width: 50%;
    margin-bottom: 20px;
  }

  .modal-header {
    font-size: 24px;
  }

  .modal-sub-header {
    font-size: 20px;
  }

  .modal-text-right {
    text-align: right;
  }
</style>

<template>
  <div v-if="show && request" class="container">
      <div class="modal-container">
        <div class="modal-half-row modal-header">Request Info:</div>
        <div class="modal-half-row modal-header modal-text-right"><div v-on:click="hideModal">X</div></div>

        <div class="modal-half-row">Method: {{request.method}}</div>
        <div class="modal-half-row modal-text-right">Url:{{request.url}}</div>

        <div class="modal-row modal-sub-header">Headers</div>
        <div class="modal-row" v-for="header in request.headers.split('|::|')">{{header}}</div>

        <div class="modal-row modal-sub-header">Query Params</div>
        <div class="modal-row">{{request.query_params}}</div>

        <div class="modal-row modal-sub-header">Body</div>
        <div class="modal-row">{{request.body}}</div>
      </div>
  </div>
</template>

<script>
  export default {
      props: ['show', 'request'],
      methods: {
        hideModal() {
          this.$props.show = false;
        }
      }
  }
</script>

<style>
</style>
