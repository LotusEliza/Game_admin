<template>
        <b-container fluid class="players">
            <h2 v-if="$mq === 'md'">Players hello</h2>
            <b-row>
<!-- ********************************** FILTER *************************************** -->
                <b-col lg="6" class="my-1">
                    <b-form-group
                            label="Filter"
                            label-cols-sm="3"
                            label-align-sm="right"
                            label-size="sm"
                            label-for="filterInput"
                            class="mb-0"
                    >
                        <b-input-group size="sm">
                            <b-form-input
                                    v-model="filter"
                                    type="search"
                                    id="filterInput"
                                    placeholder="Type to Search"
                            ></b-form-input>
                            <b-input-group-append>
                                <b-button :disabled="!filter" @click="filter = ''">Clear</b-button>
                            </b-input-group-append>
                        </b-input-group>
                    </b-form-group>
                </b-col>

                <b-col lg="6" class="my-1">
                    <b-form-group
                            label="Filter On"
                            label-cols-sm="3"
                            label-align-sm="right"
                            label-size="sm"
                            description="Leave all unchecked to filter on all data"
                            class="mb-0">
                        <b-form-checkbox-group v-model="filterOn" class="mt-1">
                            <b-form-checkbox value="Tg">Tg</b-form-checkbox>
                        </b-form-checkbox-group>
                    </b-form-group>
                </b-col>
<!--**************************** PAGINATION **********************************-->
                <b-col class="my-1" md="6" offset-md="3">
                    <b-pagination
                        align="fill"
                        size="sm"
                        class="my-0"
                        :total-rows="totalItems"
                        v-model="currentPage"
                        :per-page="perPage"
                    ></b-pagination>
                </b-col>
            </b-row>
<!-- ************************** TABLE ************************** -->
            <b-row class='p-3'>
                <b-table 
                    small
                    stacked="md"
                    show-empty 
                    :items="items"
                    :fields="fields" 
                    :current-page="currentPage" 
                    :per-page="0"
                    :filter="filter"
                    :filterIncludedFields="filterOn"
                    @filtered="onFiltered"
                >
                    <template v-slot:cell(Profile)="row">
                            <router-link
                            :to="{ name: 'profile', params: {item: row.item} }"
                            tag="button" class="btn btn-secondary btn-sm btn-block" append>
                                Profile
                            </router-link>
                    </template>
                </b-table>
            </b-row>
        </b-container>
</template>

<script>
    export default {
       data() {
    return {
      items: [],
      fields: [
                { key: 'Tg', label: 'Tg'},
                { key: 'Chat', label: 'Chat'},
                { key: 'Faction', label: 'Faction'},
                { key: 'Story', label: 'Story'},
                { key: 'Name', label: 'Name'},
                { key: 'Registered', label: 'Registered'},
                { key: 'LastActive', label: 'LastActive'},
                { key: 'Referrer', label: 'Referrer'},
                { key: 'Location', label: 'Location'},
                { key: 'MaxHP', label: 'MaxHP'},
                { key: 'HP', label: 'HP'},
                { key: 'EquipW', label: 'EquipW'},
                { key: 'EquipA', label: 'EquipA'},
                { key: 'PosX', label: 'PosX'},
                { key: 'PosY', label: 'PosY'},
                { key: 'Air', label: 'Air'},
                { key: 'Profile', label: ''},
               ],
      currentPage: 1,
      perPage: 10,
      totalItems: 0,
      filter: null,
      filterOn: [],   
    }
  },
  mounted() {
    this.fetchData().catch(error => {
      console.error(error)
    })
  },
  methods: {
    async fetchData(){
        window.axios.get(`/players?page=${this.currentPage}`).then((resp) => {
            this.items = resp.data.Players
            this.totalItems = resp.data.TotalPlayers
            console.log()
        }).catch(function (err) {
            console.log(err);
        });
    },
    onFiltered(filteredItems) {
        // Trigger pagination to update the number of buttons/pages due to filtering
        this.totalRows = filteredItems.length
    },
  },
  watch: {
    currentPage: {
      handler: function(value) {
        this.fetchData().catch(error => {
          console.error(error)
        })
      }
    }
  },

    }
</script>
<style >
</style>