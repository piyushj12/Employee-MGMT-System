<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-sm-6">
        <div class="row">
          <div class="col-sm-6">
            <select class="form-control" v-model="field">
              <option disabled value="">Please Select Field</option>
              <option
                v-for="(field, index) in fields"
                :key="index"
                class="form-control"
                :value="field.replace(/\s+/g, '')"
              >
                {{ getTitleCase(field) }}
              </option>
            </select>
          </div>
          <div class="col-sm-6">
            <div class="input-group">
              <div class="form-outline">
                <input
                  type="text"
                  id="form1"
                  placeholder="Search Employee"
                  class="form-control"
                  v-model="search"
                />
              </div>
              <button
                type="button"
                @click="searchEmployee()"
                class="btn btn-primary"
              >
                <i class="fas fa-search"></i>
              </button>
            </div>
          </div>
        </div>
      </div>
      <div class="col-sm-6">
        <div class="text-end">
          <button class="btn btn-primary" @click="getAllemployees('active')">
            Active Employees</button
          >&nbsp;
          <button
            class="btn btn-secondary"
            @click="getAllemployees('inactive')"
          >
            Inctive Employees
          </button>
        </div>
      </div>
    </div>
    <b-modal ref="modal-view" :title="modaltitle" @hide="resetdata()">
      <div class="p-3 d-flex flex-column justify-content-center">
        <div class="form-group">
          <div v-if="modal == 'delete' || modal == 'permanent'">
            <div class="text-center">
              <i class="fa fa-trash text-danger fa-3x" aria-hidden="true"></i>
            </div>
            <h6 class="mt-2">
              Are you sure you want to delete the following employee?
            </h6>
            <p class="text-center">
              {{ view.firstname + " " + view.lastname }}
            </p>
          </div>
          <div v-else-if="modal == 'restore'">
            <div class="text-center">
              <i class="fa fa-trash text-success fa-3x" aria-hidden="true"></i>
            </div>
            <h6 class="mt-2">
              Are you sure you want to restore the following employee?
            </h6>
            <p class="text-center">
              {{ view.firstname + " " + view.lastname }}
            </p>
          </div>
          <div v-else>
            <table class="table table-bordered table-striped">
              <tbody>
                <tr>
                  <th>First Name</th>
                  <td>
                    <p v-if="modal == 'view'">{{ view.firstname }}</p>
                    <input
                      v-else
                      v-model="view.firstname"
                      class="form-control"
                    />
                  </td>
                </tr>
                <tr>
                  <th>Last Name</th>
                  <td>
                    <p v-if="modal == 'view'">{{ view.lastname }}</p>
                    <input
                      v-else
                      v-model="view.lastname"
                      class="form-control"
                    />
                  </td>
                </tr>
                <tr>
                  <th>Gender</th>
                  <td>
                    {{ view.gender }}
                  </td>
                </tr>
                <tr>
                  <th>Department</th>
                  <td>
                    <p v-if="modal == 'view'">{{ view.department }}</p>
                    <select
                      v-else
                      v-model="view.department"
                      class="form-control"
                    >
                      <option disabled value="">Please Select Department</option>
                    <option v-for="(value,index) in departments" :key="index" :value="value">
                            {{value}}
                    </option>
                    </select>
                  </td>
                </tr>
                <tr>
                  <th>Email</th>
                  <td>
                    {{ view.email }}
                  </td>
                </tr>
                <tr>
                  <th>Yearly Holidays</th>
                  <td>
                    <p v-if="modal == 'view'">{{ view.holidays }}</p>
                    <select  v-else
                      v-model.number="view.holidays" class="form-control">
                     <option disabled value="">Please Select Holidays</option>
                    <option v-for="(holiday,index) in holidaysArr" :key="index" :value="holiday">
                            {{holiday}}
                    </option>
                </select>
                  </td>
                </tr>
                <tr>
                  <th>Remaining Holidays</th>
                  <td>
                    <p v-if="modal == 'view'">{{ view.remholidays }}</p>
                    <input
                      v-else
                      v-model.number="view.remholidays"
                      class="form-control"
                    />
                  </td>
                </tr>
                <tr>
                  <th>Manager</th>
                  <td>
                    <p v-if="modal == 'view'">{{ view.manageremail }}</p>
                    <input
                      v-else
                      v-model="view.manageremail"
                      class="form-control"
                    />
                  </td>
                </tr>
                <tr>
                  <th>Position</th>
                  <td>
                    <p v-if="modal == 'view'">{{ view.level }}</p>
                    <select v-else v-model="view.level" class="form-control">
                     <option disabled value="">Please Select Position</option>
                    <option v-for="(value,index) in levels" :key="index" :value="value">
                            {{value}}
                    </option>
                    </select>
                  </td>
                </tr>
                <tr>
                  <th>Role</th>
                  <td>
                    <p v-if="modal == 'view'">{{ view.role }}</p>
                    <select v-else v-model="view.role" class="form-control">
                      <option disabled value="">Please Select Role</option>
                      <option value="hr">HR</option>
                      <option value="employee">Employee</option>
                    </select>
                  </td>
                </tr>
                <tr>
                  <th>Address</th>
                  <td>
                    <p v-if="modal == 'view'">{{ view.address }}</p>
                    <input v-else v-model="view.address" class="form-control" />
                  </td>
                </tr>
                <tr>
                  <th>Contact</th>
                  <td>
                    <p v-if="modal == 'view'">{{ view.contact }}</p>
                    <input v-else v-model="view.contact" class="form-control" />
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <template #modal-footer="{ cancel }">
        <b-button size="md" variant="secondary" @click="cancel()">
          Cancel
        </b-button>
        <button
          v-if="modal == 'edit'"
          type="button"
          class="btn btn-primary"
          @click="EditEmployee(view)"
        >
          Edit
        </button>
        <button
          v-if="modal == 'delete' || modal == 'permanent'"
          type="button"
          class="btn btn-danger"
          @click="deleteEmployee(view)"
        >
          Delete
        </button>
        <button
          v-if="modal == 'restore'"
          type="button"
          class="btn btn-success"
          @click="restoreEmployee(view)"
        >
          Restore
        </button>
      </template>
    </b-modal>
    <div class="row mt-3">
      <div class="col-sm-12" v-if="list.length > 0">
        <table class="table table-bordered table-striped">
          <thead>
            <tr>
              <th>Sr. No.</th>
              <th>Employee Name</th>
              <th>Status</th>
              <th>Email</th>
              <th>Contact</th>
              <th>Department</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(emp, index) in list" :key="index">
              <td>{{ index + 1 }}</td>
              <td>{{ emp.firstname + " " + emp.lastname }}</td>
              <td>{{ emp.empstatus }}</td>
              <td>{{ emp.email }}</td>
              <td>{{ emp.contact }}</td>
              <td>{{ emp.department }}</td>
              <td>
                <div v-if="emp.empstatus == 'active'">
                  <button
                    class="btn btn-info btn-sm"
                    @click="viewDetails(emp, 'view')"
                  >
                    View</button
                  >&nbsp;
                  <button
                    class="btn btn-secondary btn-sm"
                    @click="EditEmployeeview(emp, 'edit')"
                  >
                    Edit</button
                  >&nbsp;
                  <button
                    class="btn btn-danger btn-sm"
                    @click="DeleteEmpview(emp, 'delete')"
                  >
                    Delete
                  </button>
                </div>
                <div v-else>
                  <button
                    class="btn btn-success btn-sm"
                    @click="viewRestore(emp, 'restore')"
                  >
                    Restore</button
                  >&nbsp;
                  <button
                    class="btn btn-danger btn-sm"
                    @click="deleteEmpPermanentlyView(emp, 'permanent')"
                  >
                    Delete Permanently</button
                  >&nbsp;
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <NoData v-else />
    </div>
  </div>
</template>

<script>
import NoData from "@/components/noDataTemplate.vue";
import UserServices from "@/services/userService.js";
import Services from "@/services/EmployeeServices.js";
export default {
  components: {
    NoData,
  },
  data() {
    return {
      list: [],
      view: {},
      modal: "",
      modaltitle: "",
      fields: ["first name", "last name", "department", "email"],
      search: "",
      field: "",
      departments:[],
      levels:[],
      holidaysArr:[]
    };
  },
  async created() {
    this.getAllemployees("active");
    this.getCompanyData()
  },
  methods: {
     getCompanyData(){
             UserServices.getCompanyData().then((data) => {
                  this.departments=data.departments
                  this.levels=data.levels
                  this.holidaysArr=data.holidays.sort()
            });
        },
    searchEmployee() {
      if (this.field == "") {
        this.$toast.open({
          message: "Select field first",
          type: "error",
          position: "top",
        });
      } else {
        this.list = this.list.filter((obj) => {
          return obj[this.field].includes(this.search);
        });
      }
    },
    async getAllemployees(status) {
      await UserServices.getAllEmployeesService({ empstatus: status }).then(
        (data) => {
          if (data != null) {
            this.list = data;
          } else {
            this.list = [];
          }
        }
      );
    },
    viewDetails(emp, modal) {
      this.modal = modal;
      this.view = emp;
      this.modaltitle = "Employee Details";
      this.$refs["modal-view"].show();
    },
    EditEmployeeview(emp, modal) {
      this.modal = modal;
      this.view = emp;
      this.modaltitle = "Edit Employee Details";
      this.$refs["modal-view"].show();
    },
    async EditEmployee(view) {
      await Services.editProfile(view).then(() => {
        this.getAllemployees("active");
        this.$refs["modal-view"].hide();
      });
    },
    DeleteEmpview(emp, modal) {
      this.modal = modal;
      this.view = emp;
      this.modaltitle = "Delete Employee";
      this.$refs["modal-view"].show();
    },
    deleteEmpPermanentlyView(emp, modal) {
      this.modal = modal;
      this.view = emp;
      this.modaltitle = "Delete Employee Permanently";
      this.$refs["modal-view"].show();
    },
    viewRestore(emp, modal) {
      this.modal = modal;
      this.view = emp;
      this.modaltitle = "Restore Employee";
      this.$refs["modal-view"].show();
    },
    async deleteEmployee(view) {
      if (this.modal == "delete") {
        await Services.editProfile({
          email: view.email,
          empstatus: "inactive",
        }).then(() => {
          this.getAllemployees("active");
          this.$refs["modal-view"].hide();
        });
      } else if (this.modal == "permanent") {
        await Services.deleteEmpPermanently({
          email: view.email,
          empid: view.empid,
        }).then(() => {
          this.getAllemployees("active");
          this.$refs["modal-view"].hide();
        });
      }
    },

    async restoreEmployee(view) {
      await Services.editProfile({
        email: view.email,
        empstatus: "active",
      }).then(() => {
        this.getAllemployees("inactive");
        this.$refs["modal-view"].hide();
      });
    },
    resetdata() {
      if (
        this.modal == "view" ||
        this.modal == "edit" ||
        this.modal == "delete"
      ) {
        this.getAllemployees("active");
      } else {
        this.getAllemployees("inactive");
      }
    },
  },
};
</script>

<style></style>
