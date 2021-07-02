<template>
  <v-app>
    <v-app-bar app class="appbar" dark>
      <v-app-bar-title class="title">Go KMeans Table</v-app-bar-title>
    </v-app-bar>
    <v-main>
      <v-container>
        <v-row justify="center">
          <v-col cols="12" sm="6" md="4"><v-text-field v-model="clusters" label="K Clusters to process" dark color="#da71ff" outlined type="number"/></v-col>
          <v-col cols="6" sm="3" md="2"><v-btn color="white" rounded="xl" text @click="load" class="load" elevation="3" large>Load</v-btn></v-col>
        </v-row>
      </v-container>
      <v-data-table
          dark
          :headers="headers"
          :items="dataset['data']"
          :items-per-page="25"
          :footer-props="{
            'items-per-page-options': [15, 25, 50, 75]
          }"
          class="elevation-5 table"
      ></v-data-table>
    </v-main>
  </v-app>
</template>

<script>
import axios from "axios";
export default {
  name: 'App',
  components: {
  },
  data() {
    return {
      headers: [
        {text: 'Id', value: 'id'},
        {text: 'Estrato', value: 'estrato_socioeconomico'},
        {text: 'Seguridad Nocturna', value: 'seguridad_nocturna'},
        {text: 'Grupo Edad', value: 'grupos_edad'},
        {text: 'Confianza Policia', value: 'confianza_policia'},
        {text: 'Pronto Delito', value: 'pronto_delito'},
        {text: 'Cluster', value: 'cluster'},
      ],
      dataset: '',
      clusters: 4,
    }
  },
  mounted () {
    axios
        .get('http://localhost:3000/4')
        .then(response => (this.dataset = response))
  },
  methods: {
    load() {
      axios
          .get('http://localhost:3000/'+this.clusters)
          .then(response => (this.dataset = response))
    }
  }
}
</script>

<style lang="scss">
/*scrollbar*/
/* width */
::-webkit-scrollbar {
  width: 5px;
}

/* Track */
::-webkit-scrollbar-track {
  background: #222222;
}

/* Handle */
::-webkit-scrollbar-thumb {
  background: #da71ff;
  border-radius: 3px !important;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: #eeeeee;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 30px;
  background-color: #252525;
}
.appbar {
  background: linear-gradient(-90deg, #634dfc 0%, #da71ff 100%);
}
.load{
  transition: 3s;
  inset: 5px;
  width: 6vw;
  background: linear-gradient(-90deg, #634dfc 0%, #da71ff 100%);
}
.table{
  margin-bottom: 30px;
}
.v-data-table > .v-data-table__wrapper > table > tbody > tr > th,
.v-data-table > .v-data-table__wrapper > table > thead > tr > th,
.v-data-table > .v-data-table__wrapper > table > tfoot > tr > th {
  font-size: 16px !important;
}
</style>
