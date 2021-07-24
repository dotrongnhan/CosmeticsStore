<template>
  <div class="flex-w mb-4 d-flex justify-content-end">
    <div>
      <Pagination
          :length="allProducts.length"
          :pageSize="limit"
          :pageIndex="pageIndex"
          @change="changePage"
      />
    </div>
  </div>
<table class="table table-striped">
    <thead>
    <tr>
      <th scope="col" style="width: 5%;" class="t-center col-1"/>
      <th scope="col" class="t-center col-2">Username</th>
      <th scope="col" class="t-center col-2">Product Name</th>
      <th scope="col" class="t-center col-1">Quantity</th>
      <th scope="col" class="t-center col-3">Address Shipping</th>
      <th scope="col" style="width: 18%;" class="t-center col-2">Status</th>
      <th scope="col" class="t-center col-2">Action</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="(allProduct, index) in allProducts" :key="allProduct.id">
      <th scope="row" class="t-center parentCenter"><div class="center">{{index + 1}}</div></th>
      <td class="t-center parentCenter"><div class="center">{{allProduct.user.id === 0 ? "Visitor" : allProduct.user.full_name}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{allProduct.product.product_name}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{allProduct.quantity}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{allProduct.user.id === 0 ? "Default Address" : allProduct.user.address}}</div></td>      
      <td class="parentCenter">
        <div class="row center">
          <div class="col-6 flex-c">
            <button type="button" class="btn btn-warning" v-show="allProduct.is_paid === 0">Not Paid</button>
            <button type="button" class="btn btn-success" v-show="allProduct.is_paid === 1">Paid</button>
          </div>
          <div class="col-6 flex-c">
            <button type="button" class="btn btn-warning" v-show="allProduct.ship === 'Preparing'"> {{allProduct.ship}} </button>
            <button type="button" class="btn btn-primary" v-show="allProduct.ship === 'Shipping'"> {{allProduct.ship}} </button>
            <button type="button" class="btn btn-success" v-show="allProduct.ship === 'Shipped'"> {{allProduct.ship}} </button>
          </div>
        </div>
      </td>
      <td class="t-center parentCenter">
          <button type="button" class="btn btn-danger" @click="deleteOrder(allProduct.id)">Delete</button>
      </td>    
    </tr>
    </tbody>
  </table>
</template>

<script>
import {mapActions, mapState} from "vuex";
import Pagination from "./Pagination";

export default {
    name: "ManageOrders",
    components: {
      Pagination
    },
    data() {       
    return {
      pageIndex: 1,
      limit: 10,
    }
    },
    computed: {
    ...mapState("cart", ["allProducts"]),
    },
    created() { 
        this.getAllOrders({});
    },
    methods: {
    ...mapActions("cart", ["getAllOrders", "deleteOrder"]),
    changePage(value) {
      this.pageIndex = value
      this.getAllOrders({offset: value});
    }
    }
}
</script>