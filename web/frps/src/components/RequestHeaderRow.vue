<style>
.header-row {
    width: 100%;

    display: flex;
}

.key-row {
    width: 35%;
    padding-right: 20px;
}

.val-row {
    width: 65%;
}

.header-input {
    width: 100%;
}
</style>

<template>
  <div class="header-row">
      <div class="key-row">
          <input class="header-input" v-model="headerKey" />
      </div>
      <div class="val-row">
          <input class="header-input" v-model="headerVal" />
      </div>
  </div>
</template>

<script>
  export default {
    props: [ 'header' ],
    computed: {
        headerKey() {
            return this.header.split('|:::|')[0].trim();
        },
        headerVal() {
            return this.header.split('|:::|')[1].trim();
        },
    }
  }
</script>

<style>
</style>
