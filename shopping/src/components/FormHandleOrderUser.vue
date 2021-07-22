<template>
  <div id="myModal" class="modal">
    <div class="modal-content" >
        <div class="container-table-cart pos-relative">
              <div class="wrap-table-shopping-cart bgwhite">
                <table class="table-shopping-cart" v-if="customerProducts != null">
                  <tr class="table-head">
                    <th class="column-1"></th>
                    <th class="column-2">Product</th>
                    <th class="column-3">Price</th>
                    <th class="column-4 t-center">Quantity</th>
                    <th class="column-4 t-center">Total</th>
                  </tr>
                  <tr
                      class="table-row"
                      v-for="product in customerProducts"
                      :key="product.id"
                  >
                    <td class="column-1">
                      <div class="cart-img-product b-rad-4 o-f-hidden">
                        <img :src="product.product.image" alt="IMG-PRODUCT" />
                      </div>
                    </td>
                    <td class="column-2">{{ product.product.product_name }}</td>
                    <td class="column-3">${{ product.product.price }}</td>
                    <td class="column-4">
                      {{product.quantity}}
                    </td>
                    <td class="column-5 t-center">${{ product.quantity * Number(product.product.price) }}</td>
                  </tr>
                  <tr class="table-row">
                    <td class="column-1" colspan="3"></td>
                    <td class="text-right p-r-10"><b>Total:</b></td>
                    <td class="column-5">$ {{ subTotal1 }}</td>
                  </tr>
                </table>
                <div v-else style="margin-bottom: 50px;">
                  There is no products in this customer's cart
                </div>
              </div>
      <div class="button">
        <button class="close-btn" @click="closeForm">Close</button>
      </div>
            </div>
    </div> 
  </div>
</template>

<script>
import { mapState, mapGetters} from "vuex";
import { currency } from "@/utils/currency";
// import { Form, Field, ErrorMessage } from "vee-validate";

export default {
  name: "FormHandleOrderUser",
  // components:{
  //   Form,
  //   Field,
  //   ErrorMessage
  // },
  props : ['customerProducts'],
  data() {
    return {
    }
  },
  methods: {
    closeForm() {
      this.$store.commit("cart/displayForm");
    }
  },
  computed: {
    currency,
    ...mapGetters("cart", ["subTotal1"]),
    ...mapState("cart", ["isDisplayFormOrder"]),
  }
}
</script>

<style scoped>
.button {
  display: flex;
  justify-content: space-around;
}
.close-btn:hover {
    color: red;
}
[type=button]:not(:disabled), [type=reset]:not(:disabled), [type=submit]:not(:disabled), button:not(:disabled) {
  border-radius: 10px;
  cursor: pointer;
  height: 40px;
  width: 400px;
}
.profile-img{
  text-align: center;
}
.profile-img img{
  width: 350px;
  height: 300px;
}
.primary-info {
  display: grid;
  grid-template-columns: auto auto;
  grid-gap: 200px;
}

.modal {
  text-align: center;
  font-size: 25px;
  display: block;
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0,0,0);
  background-color: rgba(0,0,0,0.4);
}

.modal-content {
  background-color: #fefefe;
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
}

/* The Close Button */
.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>