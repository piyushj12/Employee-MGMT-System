<template>
  <div class="container-fluid">
    <div class="" v-if="appliedleaves.length > 0">
      <table class="table table-bordered table-striped">
        <thead>
          <tr>
            <th>Applied on</th>
            <th>Manager</th>
            <th>From</th>
            <th>To</th>
            <th>No. of days</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(leaves, index) in appliedleaves" :key="index">
            <td>{{ new Date(leaves.applieddate).toLocaleString() }}</td>
            <td>{{ leaves.manageremail }}</td>
            <td>{{ new Date(leaves.fromdate).toLocaleDateString() }}</td>
            <td>{{ new Date(leaves.todate).toLocaleDateString() }}</td>
            <td>{{ leaves.numdays }}</td>
            <td>{{ leaves.status }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <NoData v-else />
  </div>
</template>
<script>
import NoData from "@/components/noDataTemplate.vue";
import Services from "../services/EmployeeServices.js";
export default {
  components: { NoData },
  props: {
    email: String,
    status: String,
    field: String,
  },
  data() {
    return {
      appliedleaves: [],
    };
  },
  created() {
    this.getEmployeeAppliedLeaves();
  },
  methods: {
    async getEmployeeAppliedLeaves() {
      await Services.getEmployeeAppliedLeaves({
        email: this.email,
        status: this.status,
        field: this.field,
      }).then((data) => {
        if (data != null) {
          this.appliedleaves = data;
        } else {
          this.appliedleaves = [];
        }
      });
    },
  },
};
</script>
