(function(Vue){
    var app = new Vue({
        el: 'body',

        data: {
            tasks: [],
            newTask: {}
        },

        created: function() {
            // load all the tasks
            // this.$http.get('/tasks').then(function(res) {
            //     this.tasks = res.data.iterms ? res.data.iterms : [];
            // });
        },

        methods: {
            createTask: function() {
                if (!$.trim(this.newTask.name)) {
                    this.newTask = {};
                    return;
                };

                this.newTask.done = false;

                // send http post req
                // this.$http.post('url', this.newTask).success(function(res) {
                //     this.newTask.id = res.created;
                //     this.tasks.push(this.newTask);

                //     this.newTask = {};
                // }).error(function(err) {
                //     console.log(err);
                // });

                this.newTask.id = '0';
                this.tasks.push(this.newTask);
                this.newTask = {};
            },

            deleteTask: function(index) {
                // send http delete req to delete task
                // this.$http.delete('/task/'+index).success(function(res) {
                //     this.$http.get('/tasks').then(function(res) {
                //          this.tasks = res.data.iterms ? res.data.iterms : [];
                //      })
                // }).error(function(err) {
                //     console.log(err);
                // });
            },

            updateTask: function(task, completed) {
                if (completed) {
                    task.done = true;
                }

                // send http put req to update task info
                // this.$http.put('/task', task).success(function(res) {
                //     this.$http.get('/tasks').then(function(res) {
                //         this.tasks = res.data.iterms ? res.data.iterms : [];
                //     });
                // }).error(function(err) {
                //     console.log(err)
                // });
            }
        }
    });
})(Vue);