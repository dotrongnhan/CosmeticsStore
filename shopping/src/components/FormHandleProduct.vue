<template>
  <div id="myModal" class="modal">
    <div class="modal-content">
      <h1>PRODUCT</h1>
        <div class="profile-img">
          <img :src="currentProduct.image" alt=""/>
        </div>
      <div style="justify-content: center;display: flex" class="container">
        <Form @submit.prevent="submit"  style="min-width: 65%;" :validation-schema="schema">
          <div  class="bo4 of-hidden size15 m-b-5">
            <Field
                v-model="currentProduct.name"
                name="name"
                type="text"
                placeholder="Product's name"
                class="sizefull s-text7 p-l-22 p-r-22"
            />
          </div>
          <ErrorMessage name="name" class="fs-18 text-danger d-block" />

          <div class="bo4 of-hidden size15 m-b-5">
            <Field
                v-model="currentProduct.image"
                name="image"
                type="text"
                placeholder="Product's image"
                class="sizefull s-text7 p-l-22 p-r-22"
            />
          </div>
          <ErrorMessage name="image" class="fs-18 text-danger d-block" />
          <div class="row">
            <div class="col-6">
              <div class="bo4 of-hidden size15 m-b-5">
                <Field
                    v-model="currentProduct.price"
                    name="price"
                    type="number"
                    placeholder="Product's price"
                    class="sizefull s-text7 p-l-22 p-r-22"
                />
              </div>
              <ErrorMessage name="price" class="fs-18 text-danger d-block" />

            </div>
            <div class="col-6">
              <div class="bo4 of-hidden size15 m-b-5">
                <Field
                    v-model="currentProduct.numberAvailable"
                    name="numberAvailable"
                    type="number"
                    placeholder="Product's number available"
                    class="sizefull s-text7 p-l-22 p-r-22"
                />
              </div>
              <ErrorMessage name="numberAvailable" class="fs-18 text-danger d-block" />

            </div>
          </div>
          <div class="row">
            <div class="col-4">
              <Select2
                  :options="[
                      { value: 'Default', label: 'Default category' },
                      { value: 'Toner', label: 'Toner' },
                      { value: 'Mask', label: 'Mask' },
                      { value: 'Lipstick', label: 'Lipstick' },
                      { value: 'Face Scream', label: 'Face Scream' },
                      { value: 'Serum', label: 'Serum' },
                    ]"
                  @change="categoryUpdate"
                  :value="currentProduct.currentCategory"
              />
            </div><div class="col-4">
              <Select2
                  :options="[
                      { value: '1', label: 'SK-II' },
                      { value: '2', label: 'MAC' },
                    ]"
                  @change="brandUpdate"
                  :value="currentProduct.currentBrand"/>
            </div><div class="col-4">
              <Select2
                  :options="[
                      { value: 0, label: 'Default status' },
                      { value: 1, label: 'Hot' },
                    ]"
                  @change="statusUpdate"
                  :value="currentProduct.currentStatus"/>
            </div>
          </div>
        </Form>
      </div>
      <div style="margin-top: 20px" class="button row m-t-30" >
        <div class="col-6">
          <button CLASS="btn btn-outline-secondary w-size2" @click="noConfirm">Back</button>
        </div>
        <div class="col-6">
          <button class="btn btn-outline-success w-size2" type="submit" @click="yesConfirm">Submit</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";
import Select2 from "./Select2";

export default {
  name: "FormHandleProduct",
  components:{
    Form,
    Field,
    ErrorMessage,
    Select2
  },
  computed: {

  },
  mounted() {

      if (this.product.category_id === 1) {
        this.currentProduct.currentCategory = { value: 1, label: 'Face Scream' }
      } else if (this.product.category_id === 2) {
        this.currentProduct.currentCategory =  { value: 1, label: 'Serum' }
      } else if (this.product.category_id === 3) {
        this.currentProduct.currentCategory =  { value: 3, label: 'Toner' }
      } else if (this.product.category_id === 4) {
        this.currentProduct.currentCategory =  { value: 4, label: 'Mask' }
      } else if (this.product.category_id === 5) {
        this.currentProduct.currentCategory =  { value: 5, label: 'Lipstick' }
      } else {
        this.currentProduct.currentCategory =  undefined
      }

      if (this.product.brand_id === 1) {
        this.currentProduct.currentBrand = { value: '1', label: 'SK-II' }
      } else if (this.product.brand_id === 2 ) {
        this.currentProduct.currentBrand = { value: '2', label: 'MAC' }
      } else {
        this.currentProduct.currentBrand = undefined
      }

      if(this.product.is_hot) {
        this.currentProduct.currentStatus = { value: 1, label: 'Hot' }
      } else {
        this.currentProduct.currentStatus = { value: 0, label: 'Default status' }

      }
  },
  props: ['product'],
  data() {
    const schema = yup.object().shape({
      name: yup
          .string()
          .required("Product's name is required"),
      image: yup
          .string()
          .required("Product's image is required"),
      price: yup
          .number()
          .required("Product's price is required"),
      numberAvailable: yup
          .number()
          .required("Product's number available is required")
    })
    return {
      schema,
      currentProduct: {
        name: this.product.product_name || '',
        price: this.product.price || 0 ,
        image: this.product.image || '',
        numberAvailable: this.product.number_available || 0,
        currentCategory: { value: 'Default', label: 'Default category' },
        currentBrand: { value: 1 , label: 'SK-II' },
        currentStatus: { value: this.product.isHot, label: 'Default status' },
      }
    }
  },
  methods: {
    noConfirm: function () {
      this.$emit("updateProduct", false)
    },
    yesConfirm: function () {
      this.$emit("confirmDeleteItem", true)
    },
    categoryUpdate(value) {
      this.currentProduct.currentCategory = value
    },
    brandUpdate(value) {
      this.currentBrand.currentBrand = value
    },
    statusUpdate(value) {
      this.currentStatus.currentStatus = value
    }

  }
}
</script>

<style scoped>
.button {
  display: flex;
  justify-content: space-around;
}
[type=button]:not(:disabled), [type=reset]:not(:disabled), [type=submit]:not(:disabled), button:not(:disabled) {
  border-radius: 10px;
  cursor: pointer;
  height: 40px;
  width: 400px;
}
.modal {
  text-align: center;
  font-size: 20px;
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
.bo4 {
  border: 0.5px solid #aaaaaa;
  border-radius: 5px;
}
</style>