<template>
   <b-container fluid class="equip">
<!-- ************************** TABLE ************************** -->
            <b-row class="p-3" v-if="$mq === 'sm'">
                 <router-link  
                            :to="{ name: 'addequip'}"
                            tag="button" class="btn btn-secondary btn-sm btn-block">
                                Add new equipment
                 </router-link>
            </b-row>
            <b-row class='p-3'>
                <b-table 
                    small
                    stacked="md"
                    show-empty 
                    :items="items"
                    :fields="fields" 
                >
                    <template v-slot:cell(Update)="row" >
                            <router-link  
                            :to="{ name: 'updequip', params: {exampleProp: row.item} }"
                            tag="button" class="btn btn-secondary btn-sm btn-block" append>
                                Update
                            </router-link>
                    </template>

                    <template v-slot:cell(Remove)="row">
                             <b-button class="btn btn-secondary btn-sm btn-block" @click="removeEquip(row.item)">
                                 Remove
                             </b-button>
                    </template>
                </b-table>
            </b-row>
            <mq-layout mq="md+">
                <b-row class="p-3">
                    <b-link  
                        :to="{ path: 'addequip'}"
                        tag="button" class="btn btn-secondary btn-lg" append>
                            Add new equipment
                    </b-link>
                </b-row>
            </mq-layout>
        </b-container>
</template>

<script>
    import { mapGetters } from 'vuex';
    export default {
       data() {
            return {
                id: null,
                 // items: [],
                  fields: [
                { key: 'ID', label: 'ID'},
                { key: 'Title', label: 'Title'},
                { key: 'Type', label: 'Type'},
                { key: 'SubType', label: 'SubType'},
                { key: 'BuyPrice', label: 'BuyPrice'},
                { key: 'SellPrice', label: 'SellPrice'},
                { key: 'Reputation', label: 'Reputation'},
                { key: 'Damage', label: 'Damage'},
                { key: 'Armor', label: 'Armor'},
                { key: 'Air', label: 'Air'},
                { key: 'Mine', label: 'Mine'},
                { key: 'Time', label: 'Time'},
                { key: 'SocketType', label: 'SocketType'},
                { key: 'Sockets', label: 'Sockets'},
                { key: 'Update', label: ''},
                { key: 'Remove', label: ''},
               ],
            }
    },
  // mounted() {
  //   this.fetchData().catch(error => {
  //     console.error(error)
  //   })
  // },
  created() {
            this.fetchData();
  },
  computed: {
      ...mapGetters([
          'items',
      ]),
  },
  methods: {
    fetchData(){
        this.$store.dispatch('getEquips')
    },
    removeEquip(item){
        console.log('this is ID' +item.ID)
        this.$store.dispatch('deleteEquip', item.ID)
    },
  }

}
</script>
<style scoped>
thead  {
  border-bottom: 4px solid   #a4a9ad;
}

.table th:last-child, .table th:first-child {
     border-radius:6px;
}

.table td {
    background-color: rgba(255, 250, 250, 0.55);
}

.table thead{
    color: #42a3b8!important;
}

.table.b-table.b-table-stacked-md > tbody > tr > [data-label]::before {
color: rgb(66, 163, 184);
}

table td:nth-child(1) {  
   background-color: rgba(0, 143, 253, 0.2)!important;
}
.equip table td:nth-last-child(-n+2)  {  
   text-align:center;
}
</style>

