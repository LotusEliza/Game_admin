<template>
    <div v-if="$route.meta.showNavigation">
        <b-field>
            <b-input placeholder="Search by ID..."
                     v-model='id'
                     type="search"
                     icon-pack="fas"
                     @keyup.enter.native="searchById"
                     rounded
                     size="is-small"
                     icon="search">
            </b-input>
        </b-field>
    </div>
</template>

<script>
    export default {
        data(){
            return{
                id: null,
            }
        },
        methods:{
            searchById () {
                this.$router.push({ path: '/players/profile/'+this.id });
                this.id = null;
            }
        }
    }
</script>
