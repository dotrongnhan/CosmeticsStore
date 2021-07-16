<template>
  <div class="flex-w  mb-4">
    <div class="bo4 w-size12 m-t-5 m-b-5 m-r-10">
      <Select2
          :options="[
                      { value: '', label: 'Default Sorting' },
                      { value: -1, label: 'Price: low to high' },
                      { value: 1, label: 'Price: high to low' },
                    ]"
          @change="sort"
          :value="direction"
      />
    </div>

    <div class="bo4 w-size12 m-t-5 m-b-5 m-r-10">
      <Select2
          :options="[
                      { value: '', label: 'Default category' },
                      { value: 4, label: 'Toner' },
                      { value: 2, label: 'Mask' },
                      { value: 5, label: 'Lipstick' },
                      { value: 3, label: 'Face Scream' },
                      { value: 6, label: 'Serum' },
                    ]"
          @change="category"
          :value="currentCategory"
      />
    </div>
    <input v-model="message" placeholder="Search here">
    <button type="button" class="ab-t-r mr-3 btn btn-primary">Add new</button>
  </div>
  <table class="table table-striped">
    <thead>
    <tr>
      <th scope="col" >#</th>
      <th scope="col" class="col-3">Name</th>
      <th scope="col" class="col-3">Price</th>
      <th scope="col" class="col-3">Category</th>
      <th scope="col" class="t-center col-2">Action</th>
    </tr>
    </thead>
    <tbody>
    <tr>
      <th scope="row">1</th>
      <td>Mark</td>
      <td>Otto</td>
      <td>@mdo</td>
      <td>
        <div class="row">
          <div class="col-6 flex-c">
            <button type="button" class="btn btn-success">Edit</button>
          </div>
          <div class="col-6 flex-c">
            <button type="button" class="btn btn-danger">Delete</button>
          </div>
        </div>
      </td>
    </tr>
    </tbody>
  </table>
</template>

<script>
import Select2 from "@/components/Select2.vue";

export default {
  name: "SettingsPage",
  components: {
    Select2,
  },
  data() {
    return {
      items: [
        { age: 40, first_name: 'Dickerson', last_name: 'Macdonald' },
        { age: 21, first_name: 'Larsen', last_name: 'Shaw' },
        { age: 89, first_name: 'Geneva', last_name: 'Wilson' },
        { age: 38, first_name: 'Jami', last_name: 'Carney' }
      ]
    }
  }
};
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
</style>