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
            Name: this.name
          },
        )
        .then((response) => {
          if (response.status == 200) {
            this.token = response.data.token
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