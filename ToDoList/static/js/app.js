(function(Vue){
    var app = new Vue({
        el: 'body',

        data: {
            tasks: [],
            newTask: {}
        },

        created: function() {
            //load all the tasks
            this.$http.get('/tasks/list').then(function(res) {
                this.tasks = res.data ? res.data : []
            });
        },

        methods: {
            createTask: function() {
                if (!$.trim(this.newTask.name)) {
                    this.newTask = {};
                    return;
                };

                this.newTask.done = false;

                //send http post req
                this.$http.post('/tasks/add', this.newTask).success(function(res) {
                    this.$http.get('/tasks/list').then(function(res) {
                        this.tasks = res.data ? res.data : []
                    })
                }).error(function(err) {
                    console.log(err);
                });

                this.newTask = {};

            },

            deleteTask: function(task) {
                // send http delete req to delete task
                this.$http.delete('/tasks/'+task.id).success(function(res) {
                    this.$http.get('/tasks/list').then(function(res) {
                         this.tasks = res.data ? res.data : [];
                     })
                }).error(function(err) {
                    console.log(err);
                });
            },

            updateTask: function(task, completed) {
                if (completed) {
                    task.done = true;
                }

                // send http put req to update task info
                this.$http.put('/tasks/edit', task).success(function(res) {
                    this.$http.get('/tasks/list').then(function(res) {
                        this.tasks = res.data ? res.data : [];
                    });
                }).error(function(err) {
                    console.log(err)
                });
            }
        }
    });
})(Vue);