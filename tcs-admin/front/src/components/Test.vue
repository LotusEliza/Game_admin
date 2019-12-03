<template>
    <div>
        <div v-for="norm in norms">
            {{norm.Tg}}
        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
            }
        },
        mounted() {
            this.$store.dispatch('getNorms')
        },
        computed:{
            norms(){return this.$store.state.norms}
        },
        methods:{
        }
    }
</script>