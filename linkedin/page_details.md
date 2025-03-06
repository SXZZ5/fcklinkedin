- linkedin ke most of the html tags ki jo IDs hain wo sab dynamic hain. To they cannot really be used as a reliable locator for an element. But acchi cheez ye hai ki thoda to predictable prefix region hai inki IDs ka. 
    - something like ye log extra styling ka ek suffix append karte hain id me, but prefix dhyaan se verify kar lo 4-5 jageh se to use karte ho prefix contains in Xpath expression.

### Job Search Page
- main ID ka ek container dikh rha hai. Jiske do childre hain ek job list area ka aur ek job description area.
- Listing area ka ek kaafi upar level ka div hai jiski completely dynamic id assigned hai jiska kuch nhi kar sakte. 
- game aisa hai ki jo actual job links hain wo lazy load hoti hain when that job-card comes into "view"/"focus". This means Xpath selector for the <a> tag does not work because link hai hi nahi waha pe. EmberJS op :(
- good thing is inka fixed hai ki kitni cards honi hai to a bunch of <li> tags hain jo static rehti hain. Inko use kar sakte hain. 
- So what we are doing is go to the first <li> tag. Scroll the element into view, wait for that element to be stable, yaha pe lord EmberJS data loading kar denge aur ab is <li> tag ke andar ek <a> tag with a valid link apko mil jayegi. 

### Job Description Region
- Luckily Job description (actual text part requirements, etc wala) contain karne wali jo div hai usko in log ne mast id='job-details' de rakha hai. Aur ids to unique hoti hi hain, to it should be enough.
- Job Title: Luckily jo <main> tag ke andar saara job listing + job descriptions panes hain. Usme one and only <h1> tag is the Job Title. To apan isko bhi reliably use kar sakte hain Job Title extract karne ko.
- Freshness ke liye ek span hai jisme I can count on it having the word "ago" inside of it. To iss Xpath construct kar liya. Waise to "ago" word kayi jageh aata hai page par but specifically xpath is //main//span[contains(text(), "ago")]. 
    - Yaha bhi we are lucky because exactly 1 hi span me "ago" text milega on entire page. Baaki saari jageh jahan "ago" use hua hai are all <time> tags. 
    - To isliye ye <span> se banaye xpath se collide nhi karta. 

Should be enough. Ab sirf JD 