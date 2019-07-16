const data = [
    {
        name: "Saurabh",
        open: "Opens at 6am",
        rating: "0.0",
        cuisines: "Cafe - Fast food, Burger...",
        address: "New Town District, Gurgaon",
        cft: "230 for two",
        img: "https://b.zmtcdn.com/images/logo/zomato_flat_bg_logo.svg",
    },
    {
        name: "Saurabh",
        open: "Opens at 6am",
        rating: "0.0",
        cuisines: "Cafe - Fast food, Burger...",
        address: "New Town District, Gurgaon",
        cft: "230 for two",
        img: "https://b.zmtcdn.com/images/logo/zomato_flat_bg_logo.svg",
    },
    {
        name: "Saurabh",
        open: "Opens at 6am",
        rating: "0.0",
        cuisines: "Cafe - Fast food, Burger...",
        address: "New Town District, Gurgaon",
        cft: "230 for two",
        img: "https://b.zmtcdn.com/images/logo/zomato_flat_bg_logo.svg",
    },
    {
        name: "Saurabh",
        open: "Opens at 6am",
        rating: "0.0",
        cuisines: "Cafe - Fast food, Burger...",
        address: "New Town District, Gurgaon",
        cft: "230 for two",
        img: "https://b.zmtcdn.com/images/logo/zomato_flat_bg_logo.svg",
    },
]

function shareTweet() {
    let counter = 0;
    return function () {
        if (counter == 2) {
            alert("Already shared two times");
            return;
        }
        counter++;
        alert("Shared");
    }
}

function getTweets(post) {
    return `
    <div class = "extendedcards">
    <div class = "card">
        <div class = "restrauntimage">
            <img src = "https://b.zmtcdn.com/images/logo/zomato_flat_bg_logo.svg">
        </div>
        <div class = "restarauntinfo">
            <div class = "top-info">
                <div class = "opensat">
                    <small class = "open">
                        Opens at 6am
                    </small>
                </div>
                <div class = "rating">
                    <div class = "ratingbox">
                        <small style = "font-family:Okra;">
                            0.0
                        </small>
                    </div>
                </div>
            </div>
            <div class = "restarauntname">
                <small style = "font-family:Okra;">
                    Restaraunt name
                </small>
            </div>
            <div class = "cuisines">
                <small style = "font-family:Okra;">
                    Cafe - Fast food, Burger...
                </small>
            </div>
            <div class = "address">
                <small style = "font-family:Okra;">
                    New Town District, Gurgaon
                </small>
            </div>
            <div class = "cft" >
                <small style = "font-family:Okra;">
                    230 for two
                </small>
            </div>
        </div>
    </div>
    <div class = "edit">
        <small style = "font-family:Okra;">
            Delete
        </small>
    </div>
    <br>
</div>
    `;
}

function displayTweets(post) {
    let cardHTML = '';
    const cardsParent = document.getElementsByClassName('extendedcards');
    data.forEach(function(post) {
        cardHTML += getTweets(post);
    });
    cardsParent[0].innerHTML = cardHTML;
    console.log(cardsParent);
}

displayTweets();