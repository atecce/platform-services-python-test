<template>
  <div id="app">
    <b-table :data="customers" :columns="columns"></b-table>

    <!-- TODO filter by email -->

    <!-- TODO make this prettier -->
    <input v-model="email" placeholder="Email">
    <input v-model="total" placeholder="Total">
    <button @click="order">Order</button>
  </div>
</template>

<script>

export default {
  name: 'app',
  data () {
    return {
      root: "http://localhost:7050",

      email: "",
      total: 0,

      columns: [
        {
          field: "Email",
          label: "Email"
        },
        {
          field: "RewardPoints",
          label: "Rewards"
        },
        {
          field: "RewardsTier",
          label: "Tier"
        },
        {
          field: "RewardsTierName",
          label: "Tier Name"
        },
        {
          field: "NextRewardsTier",
          label: "Next Tier"
        },
        {
          field: "NextRewardsTierName",
          label: "Next Tier Name"
        },
        {
          field: "NextRewardsTierProgress",
          label: "Next Tier Progress"
        },
      ]
    }
  },
  computed: {
    customers () {
      return this.get(this.root + "/customers")
    },
  },
  methods: {
    get (url) {
      var request = new XMLHttpRequest()
      request.open('GET', url, false)
      request.send(null)
      return JSON.parse(request.responseText)
    },
    getCustomerByEmail (email) {
      return this.get(this.root + "/customers/" + email)
    },
    order () {
      var request = new XMLHttpRequest()
      request.open('PATCH', this.root + "/order", false)
      request.send(JSON.stringify({
        email: this.email,
        total: parseFloat(this.total)
      }))
      // TODO refresh on order
    }
  }
}
</script>