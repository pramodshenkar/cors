var app = new Vue({
  el: '#app',
  data: {
    name: "",
    token: ""
  },
  methods: {
    onDashboard() {
      axios
        .post("http://localhost:5000/dashboard",
          {},
          {
            headers: {
              Authorization: this.token
            }
          }
        )
        .then((response) => {
          if (response.status == 200) {
            alert(this.token);
          }
        })
        .catch((error) => {
          console.log(error);
          alert(error.response.data.message);
        });
    },

    onLogin() {
      axios
        .post("http://localhost:5000/student/login",
          {
            username: "pramods11",
            password: "111",
          },
        )
        .then((response) => {
          if (response.status == 200) {
            this.token = response.data.result.token
            console.log(this.token)
            alert("Logging In");
          }
        })
        .catch((error) => {
          console.log(error);
          alert(error.response.data.message);
        });
    },
  }
})