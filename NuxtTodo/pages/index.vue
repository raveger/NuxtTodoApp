<template>
  <section class="container">
    <v-app>
      <v-main>
        <h3>新規登録欄</h3>
        <v-form>
          <v-container>
            <v-row>
              <v-col cols="12" sm="4">
                Name
                <v-text-field
                label="件名を入力"
                v-model="createTodoForm.name"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="6">
                Todo
                <v-text-field
                label="内容を入力"
                v-model="createTodoForm.todo"
                >
                  <template v-slot:append-outer>
                    <v-btn
                      dark
                      rounded
                      color="red"
                      v-on:click="doAdd"
                    >新規登録</v-btn>
                  </template>
                </v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-form>
        <hr>
        <h3>内容更新欄</h3>
        <v-form>
          <v-container>
            <v-row>
              <v-col cols="12" sm="1">
                ID
                <v-text-field
                  label="IDを入力"
                  v-model="updateTodoForm.id"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="4">
                Name
                <v-text-field
                  label="件名を入力"
                  v-model="updateTodoForm.name"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="7">
                Todo
                <v-text-field
                label="内容を入力"
                v-model="updateTodoForm.todo"
                >
                  <template v-slot:append-outer>
                    <v-btn
                      dark
                      rounded
                      color="blue"
                      v-on:click="doUpdate"
                    >内容を更新</v-btn>
                  </template>
                </v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-form>
        <hr>
        <v-simple-table fixed-header height="500px">
          <template v-slot:default>
            <thead>
              <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Todo</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="todo in todos" v-bind:key="todo.id">
                <td>No.{{ todo.ID }}</td>
                <td>{{ todo.Name }}</td>
                <td>{{ todo.Todo }}</td>
                <td>
                  <v-btn
                    dark
                    rounded
                    v-on:click="doRemove(todo.ID)"
                  >削除</v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-main>
    </v-app>
  </section>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'

// インターフェース
interface TodoType {
  id?: number;
  name: string;
  todo: string;
}

export default Vue.extend ({
  data() {
    return {
      todos: [] as TodoType[],
      createTodoForm: {} as TodoType,
      updateTodoForm: {} as TodoType,
    };
  },
  created() {
    this.getTODOs()
  },
  methods: {
    //一覧表示機能
    getTODOs() {
        axios.get('http://localhost:8081/todos')
        .then((response: any) => {
            this.todos = response.data as TodoType[]
            return this.todos
        }).catch((error: any) => {
            console.log(error);
        })
    },
    //新規登録機能
    doAdd() {
        axios.post('http://localhost:8081/todos', this.createTodoForm)
        .then((response: any) => {
            alert("登録完了しました。")
            this.getTODOs()
            this.createTodoForm = {} as TodoType
        }).catch((error: any) => {
            console.log(error);
        })
    },
    //更新機能
    doUpdate() {
        axios.put('http://localhost:8081/todos', this.updateTodoForm)
        .then((response: any) => {
            alert("更新が完了しました。")
            this.getTODOs()
            this.updateTodoForm = {} as TodoType
        }).catch((error: any) => {
            console.log(error);
        })
    },
    //削除機能
    doRemove(ID: number) {
        axios.delete(`http://localhost:8081/todos?id=${ID}`)
        .then((response: any) => {
            this.getTODOs()
        }).catch((error: any) => {
            console.log(error);
        })
    }
  }
})
</script>