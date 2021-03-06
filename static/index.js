let app = new Vue({
    el: '#app',
    data () {
        return {
            personName: null,
            prevName: null,
            teamName: null,
            prevTeam: null,
            teams: {},
            members: {},
        }
    },
    methods: {
        submitTeam () {
            const self = this;
            self.prevName = self.personName;

            axios
                .get("/api/teams/" + self.personName)
                .then(
                    function (response) {
                        let teamData = response.data;

                        if (teamData === null) {
                            self.teams = false;
                        }
                        else {
                            self.teams = teamData;
                        }

                    }
                )
        },

        submitMembers () {
            const self = this;
            self.prevTeam = self.teamName;

            axios
                .get("/api/members/" + this.teamName)
                .then(
                    function (response) {
                        let memberData = response.data;

                        if (memberData === null) {
                            self.members = false;
                        }
                        else {
                            self.members = memberData;
                        }
                    }
                )

        },

        isObject (value) {
            return typeof value === 'object'
        }
    }

});
