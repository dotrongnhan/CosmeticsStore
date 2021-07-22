<template>
<table class="table table-striped">
    <thead>
    <tr>
      <th scope="col" style="max-width: 3%" class="col-1"/>
      <th scope="col" class="t-center col-2">User Name</th>
      <th scope="col" class="t-center col-3">Product Name</th>
      <th scope="col" class="t-center col-2">Quantity</th>
      <th scope="col" class="t-center col-2">Status</th>
      <th scope="col" class="t-center col-2">Action</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="(allProduct, index) in allProducts" :key="allProduct.id">
      <th scope="row" class="parentCenter"><div class="center">{{index + 1}}</div></th>
      <td class="t-center parentCenter"><div class="center">{{allProduct.user.id === 0 ? "Visitor" : allProduct.user.full_name}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{allProduct.product.product_name}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{allProduct.quantity}}</div></td>  
      <td class="parentCenter">
        <div class="row center">
          <div class="col-6 flex-c">
            <button type="button" class="btn btn-primary" v-show="allProduct.is_paid === 0">Not Paid</button>
            <button type="button" class="btn btn-success" v-show="allProduct.is_paid === 1">Paid</button>
          </div>
          <div class="col-6 flex-c">
            <button type="button" class="btn btn-warning"> On Shipping </button>
          </div>
        </div>
      </td>
      <td class="t-center parentCenter">
          <button type="button" class="btn btn-danger" @click="deleteOrder(allProduct.id)">Delete</button>
      </td>    
    </tr>
    </tbody>
  </table>
    <!-- <FormHandleOrder v-show="isDisplayForm" :order="order"/> -->
</template>

<script>
import {mapActions, mapState} from "vuex";
// import FormHandleOrder from "./FormHandleOrder";

export default {
    name: "ManageOrders",
    components: {
    //   FormHandleOrder
    },
    data() {       
    return {
    //  order: {},
    }
    },
    created() {
        this.$store.dispatch("cart/getAllOrders");
    },
    computed: {
    ...mapState("cart", ["allProducts"]),
    },
    methods: {
    ...mapActions("cart", ["getAllOrders", "deleteOrder"]),
    // getOrder(order) {
    //   this.order = order
    // },
    }
}
</script>