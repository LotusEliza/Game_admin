import Players from './views/Players'
import Profile from './views/Profile'
import Equipments from './views/Equipments'
import UpdateEquip from './views/UpdateEquip'
import AddEquip from './views/AddEquip'
import Norms from './views/Norms'
import test from './views/test'

export default [
    {
        path: '/players',
        name: 'players',
        component: Players,
        meta: { showNavigation: true },
    },
    {
        path: '/players/profile/:tg',
        name: 'profile',
        component: Profile
    },
    {
        path: '/equipments/update/:tg',
        name: 'updateEquip',
        component: UpdateEquip,
    },
    {
        path: '/equipments/add',
        name: 'addEquip',
        component: AddEquip,
    },
    {
        path: '/equipments',
        name: 'equipments',
        component: Equipments,
    },
    {
        path: '/players/profile/:tg/norms',
        name: 'norms',
        component: Norms
    },
    {
        path: '/test',
        name: 'test',
        component: test
    },

]
