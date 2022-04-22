<style>
  .requests-container {
    width: 100%;
    margin: 40px;
  }

  .flex-20 {
    width: 20%;
  }

  .flex-10 {
    display: flex;
    justify-content: center;
    width: 10%;
  }

  .flex-40 {
    width: 40%;
  }

  .flex-50 {
    width: 50%;
  }

  .flex-10-s {
    width: 10%;
  }

  .request-row {
    width: 100%;

    padding: 10px;

    display: flex;
    flex-wrap: wrap;
  }

  .request-row-bottom {
    font-weight: bold;
    margin-bottom: 5px;
    border-bottom: 1px solid black;
  }

  .request-icon {
    margin-top: 7px;
  }
</style>

<template>
  <div class="requests-container">
    <div class="request-row request-row-bottom">
      <div class="flex-10-s">Method</div>
      <div class="flex-20">Timestamp</div>
      <div class="flex-50">Url</div>
      <div class="flex-10">Edit</div>
      <div class="flex-10">Replay</div>
    </div>
    <div v-for="request in requests" class="request-row">
      <div class="flex-10-s">{{request.method}}</div>
      <div class="flex-20">{{formatDate(request.created)}}</div>
      <div class="flex-50">{{request.url}}</div>
      <div class="flex-10" v-on:click="openRequestModal(request.id)">
        <i class="fa-solid fa-pen-to-square fa-xl request-icon"></i>
      </div>
      <div class="flex-10" v-on:click="replayRequest(request)">
        <i class="fa-solid fa-recycle fa-xl request-icon"></i>
      </div>
    </div>


    <RequestModal :show=showModal :request=selectedRequest />
  </div>
</template>

<script>
  import moment from 'moment';
  import RequestModal from './RequestModal.vue';

  export default {
    data() {
      return {
        requests: new Array(),
        showModal: false,
        selectedRequest: false,
      }
    },
    created() {
      this.fetchData()
    },
    watch: {
      '$route': 'fetchData'
    },
    methods: {
      formatDate(date) {
        return moment(String(date)).format('MM/DD/YYYY hh:mm A');
      },
      replayRequest(request) {
        let url = request.url;
        if (typeof request.query_params === 'string' && request.query_params.length) {
          url += `?${request.query_params}`;
        }

        // TODO: Clean this up
        let headers = {};
        let rawHeaders = request.headers.split('|::|');
        rawHeaders.forEach(raw => {
          let rawHeadersSplit = raw.split(':');
          let key = rawHeadersSplit[0].trim().slice(1, -1);
          let value = rawHeadersSplit[1].trim().slice(1, -1);

          headers[key] = value;
        });

        let options = {
          method: request.method,
          headers: new Headers(headers),
        }

        if (typeof request.body === 'string' && request.body.length) {
          options.body = request.body;
        }

        fetch(url, options)
        .then(res => {
          console.log(res);
        });
      },
      openRequestModal(id) {
        let request = this.requests.find(request => {
          return request.id === id;
        });

        if (!request) {
          return;
        }

        this.selectedRequest = request;

        this.showModal = true;
      },
      fetchData() {
        fetch('../api/requests', {credentials: 'include'})
        .then(res => {
          return res.json()
        }).then(json => {
          this.requests = json.data
        });
      }
    },
    components: {
      RequestModal,
    }
  }
</script>

