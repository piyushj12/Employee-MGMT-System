import axios from "axios";
// import router from '../router'
// import errorjs from "../Javascript/error"
const apiClient = axios.create({
  baseURL: "http://localhost:4700",
  withCredentials: false,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
    Authorization: localStorage.getItem("token"),
  },
});

export default {
  async getProfile(payload) {
    try {
      let res = await apiClient.post("/r/get/profile", payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async editProfile(payload) {
    try {
      let res = await apiClient.put("/r/edit/profile", payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async getEmployeeLeaves(payload) {
    try {
      let res = await apiClient.post("/r/get/leaves", payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async getEmployeeAppliedLeaves(payload) {
    try {
      let res = await apiClient.post("/r/get/applied/leaves", payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async submitforapproval(payload) {
    try {
      let res = await apiClient.post("/r/apply/leaves", payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async updateLeaves(payload) {
    try {
      let res = await apiClient.put("/r/update/leave/status", payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async deleteEmpPermanently(payload) {
    try {
      let res = await apiClient.delete("/r/delete/permanently/"+payload.empid);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
};
