<template>
<table class="table table-striped">
    <thead>
    <tr>
      <th scope="col" style="max-width: 3%" class="col-1"/>
      <th scope="col" class="t-center col-2">Name</th>
      <th scope="col" class="t-center col-1">Gender</th>
      <th scope="col" class="t-center col-1">Email</th>
      <th scope="col" class="t-center col-2">Phone</th>
      <th scope="col" class="t-center col-2"></th>      
      <th scope="col" class="t-center col-2">Action</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="(user, index) in users" :key="user.id">
      <th scope="row" class="parentCenter"><div class="center">{{index + 1}}</div></th>
      <td class="t-center parentCenter"><div class="center">{{user.full_name}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{user.gender}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{user.email}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{user.phone}}</div></td>
      <td class="t-center parentCenter"><button type="button" class="btn btn-primary" @click="openFormOrder(); getCartByCustomerId(user.id)">View Cart</button></td>            
      <td class="parentCenter">
        <div class="row center">
          <div class="col-6 flex-c">
            <button type="button" class="btn btn-success" @click="openForm(); getUser(user)">Show details</button>
          </div>
          <div class="col-6 flex-c">
            <button type="button" class="btn btn-danger" @click="deleteUser(user.id)">Delete</button>
          </div>
        </div>
      </td>
    </tr>
    </tbody>
  </table>
    <FormHandleOrderUser v-show="isDisplayFormOrder" :customerProducts="customerProducts"/>
    <FormHandleUser v-show="isDisplayForm" :user="user"/>
</template>

<script>
import {mapActions, mapState} from "vuex";
import FormHandleUser from "./FormHandleUser";
import FormHandleOrderUser from "./FormHandleOrderUser";

export default {
    name: "ManageUsers",
    components: {
      FormHandleUser,
      FormHandleOrderUser
    },
    data() {       
    return {
     user: {},
     cart: {},
    }
    },
    created() {
        this.$store.dispatch("users/getUsers");
    },
    computed: {
    ...mapState("users", ["users", "isDisplayForm"]),
    ...mapState("cart", ["customerProducts", "isDisplayFormOrder"]),    
    },
    methods: {
    ...mapActions("users", ["getUsers", "deleteUser"]),
    ...mapActions("cart", ["getCartByCustomerId"]),
    openForm() {
        this.$store.commit("users/displayForm");
    },
    openFormOrder() {
        this.$store.commit("cart/displayForm");
    },
    getUser(user) {
      this.user = user
    },
    }
}
</script>