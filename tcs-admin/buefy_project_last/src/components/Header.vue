<template>
    <b-navbar>
        <template slot="start">
            <b-navbar-item href="#/players" >
                Players
            </b-navbar-item>
            <b-navbar-item href="#/equipments">
                Equipments
            </b-navbar-item>
        </template>
        <template slot="end">
            <b-navbar-item tag="div">
                <slot></slot>
<!--                <div v-if="$route.meta.showNavigation">-->
<!--                    <b-field>-->
<!--                        <b-input placeholder="Search by ID..."-->
<!--                                 v-model='id'-->
<!--                                 type="search"-->
<!--                                 icon-pack="fas"-->
<!--                                 @keyup.enter.native="searchById"-->
<!--                                 rounded-->
<!--                                 size="is-small"-->
<!--                                 icon="search">-->
<!--                        </b-input>-->
<!--                    </b-field>-->
<!--                </div>-->
            </b-navbar-item>

            <b-navbar-item tag="div">
                <div class="buttons">
                    <a class="button is-primary">
                        <strong>Sign up</strong>
                    </a>
                    <a class="button is-light">
                        Log in
                    </a>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>

<script>
    export default {
        name: 'Header',
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

<style scoped>

</style>

