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
      let res = await apiClient.delete("/ah/delete/permanently/"+payload.empid);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async SearchEmployee(payload) {
    try {
      let res = await apiClient.post("/r/search",payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async DashboardData(payload) {
    try {
      let res = await apiClient.post("/r/dashboard/data",payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async CaptureClockIn(payload) {
    try {
      let res = await apiClient.post("/he/capture/clockin",payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async CaptureClockOut(payload) {
    try {
      let res = await apiClient.post("/he/capture/clockout",payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async isClockedIn(payload) {
    try {
      let res = await apiClient.post("/r/check/clockin/exists",payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async isClockedOut(payload) {
    try {
      let res = await apiClient.post("/r/check/clockout/exists",payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
  async viewAttendanceData(payload) {
    try {
      let res = await apiClient.post("/r/get/attendance",payload);
      return res.data;
    } catch (error) {
      return error.response.data;
    }
  },
};
