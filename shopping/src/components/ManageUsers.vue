<template>
<table class="table table-striped">
    <thead>
    <tr>
      <th scope="col" style="max-width: 3%" class="col-1"/>
      <th scope="col" class="t-center col-2">Name</th>
      <th scope="col" class="t-center col-1">Gender</th>
      <th scope="col" class="t-center col-1">Email</th>
      <th scope="col" class="t-center col-2">Phone</th>
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
      <td class="parentCenter">
        <div class="row center">
          <div class="col-6 flex-c">
            <button type="button" class="btn btn-success" @click="openForm">Show more</button>
          </div>
          <div class="col-6 flex-c">
            <button type="button" class="btn btn-danger">Delete</button>
          </div>
        </div>
      </td>
    </tr>
    </tbody>
  </table>
    <FormHandleUser v-show="isDisplayForm"/>
</template>

<script>
import {mapActions, mapState} from "vuex";
import FormHandleUser from "./FormHandleUser";

export default {
    name: "ManageUsers",
    components: {
      FormHandleUser
    },
    data() {       
    },
    created() {
        this.$store.dispatch("users/getUsers");
    },
    computed: {
    ...mapState("users", ["users", "isDisplayForm"]),
    },
    methods: {
    ...mapActions("users", ["getUsers"]),
    openForm() {
        this.$store.commit("users/displayForm");
    }
    }
}
</script>