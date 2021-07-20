<template>
  <div class="flex-w mb-4 d-flex justify-content-between">
    <div class="d-flex">
      <div class="w-size12 m-t-5 m-b-5 m-r-10">
        <Select2
            :options="[
                      { value: 'Default', label: 'Default Sorting' },
                      { value: 'ASC', label: 'Price: low to high' },
                      { value: 'DESC', label: 'Price: high to low' },
                    ]"
            @change="sort"
            :value="direction"
        />
      </div>

      <div class="w-size12 m-t-5 m-b-5 m-r-10">
        <Select2
            :options="[
                      { value: 'Default', label: 'Default category' },
                      { value: 'Toner', label: 'Toner' },
                      { value: 'Mask', label: 'Mask' },
                      { value: 'Lipstick', label: 'Lipstick' },
                      { value: 'Face Scream', label: 'Face Scream' },
                      { value: 'Serum', label: 'Serum' },
                    ]"
            @change="category"
            :value="currentCategory"
        />
      </div>
      <input v-model="message" placeholder="Search here">
    </div>
    <div>
      <Pagination
          :length="count"
          :pageSize="limit"
          :pageIndex="pageIndex"
          @change="changePage"
      />
    </div>
    <button @click="changeDisplay(true)" type="button" class="ab-t-r mr-3 btn btn-primary">Add new</button>
  </div>
  <table class="table table-striped">
    <thead>
    <tr>
      <th scope="col" style="max-width: 3%" class="col-1"/>
      <th scope="col" class="t-center col-1">Image</th>
      <th scope="col" class="t-center col-4">Name</th>
      <th scope="col" class="t-center col-1">Price</th>
      <th scope="col" class="t-center col-1">Category</th>
      <th scope="col" class="t-center col-1">Brand</th>
      <th scope="col" class="t-center col-2">Action</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="(product, index) in products" :key="product.id">
      <th scope="row" class="parentCenter"><div class="center">{{index + 1}}</div></th>
      <td><img style="width: 100px" :src="product.image" alt=""></td>
      <td class="t-center parentCenter"><div class="center">{{product.product_name}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{currency(product.price)}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{product.category_name}}</div></td>
      <td class="t-center parentCenter"><div class="center">{{product.brand_name}}</div></td>
      <td class="parentCenter">
        <div class="row center">
          <div class="col-6 flex-c">
            <button @click="changeDisplay(true); changeProduct(product)" type="button" class="btn btn-success">Edit</button>
          </div>
          <div class="col-6 flex-c">
            <button @click="deleteProduct(product.id)" type="button" class="btn btn-danger">Delete</button>
          </div>
        </div>
      </td>
    </tr>
    </tbody>
  </table>
  <FormHandleProduct :product="willChange" @close="changeDisplay" v-if="isDisplay" />
</template>

<script>
import Select2 from "./Select2";
import Pagination from "./Pagination";
import {mapActions, mapState} from "vuex";
import {currency} from "../utils/currency";
import FormHandleProduct from "./FormHandleProduct";

export default {
  name: "ManageProducts",
  components: {
    Select2,
    Pagination,
    FormHandleProduct
  },
  data() {
    return {
      isDisplay: false,
      value: {value: "Default"},
      direction: {value : "Default"},
      pageIndex: 1,
      limit: 16,
      willChange: {}
    }
  },
  computed: {
    ...mapState("products", ["products", "count"]),
    currentCategory() {
      return this.$store.state.products.category
    }
  },
  methods: {
    currency,
    ...mapActions("products", ["getProducts", "deleteProduct"]),
    changeProduct(product) {
      this.willChange = product
    },
    changeDisplay(value){
      this.willChange = {}
      this.isDisplay = value
    },
    sort(direction) {
      this.getProducts({category: this.value.value,sortType: direction.value, offset: 1})
      this.direction = direction
      this.pageIndex = 1
    },
    category(value) {
      this.getProducts({category: value.value, offset: this.pageIndex, sortType: this.direction.value})
      this.$store.commit('products/GET_PRODUCT_BY_CATEGORY', value)
      this.value = value
    },
    changePage(value) {
      this.pageIndex = value
      this.$store.dispatch("products/getProducts", {category: this.value.value,sortType: this.direction.value ,offset: value});
    }
  }
}
</script>

<style scoped>
input {
  box-sizing: border-box;
  cursor: pointer;
  display: block;
  background-color: #fff;
  border: 1px solid #aaa !important;
  border-radius: 4px;
  user-select: none;
  -webkit-user-select: none;
  height: 27px;
  margin-top: 5px;
  padding-left: 10px;
}
.center {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%,-50%);
}
.parentCenter {
  position: relative;
}
</style>