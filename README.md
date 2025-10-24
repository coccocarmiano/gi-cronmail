# gi-cronmail -- Birthay e-mails on a budget

## What is `gi-cronmail`

`gi-cronail` is a small app that sends birthday emails to the associates of Giovani INformazione APS ASd, a smaill NPO based in Carmiano, province of Lecce, in the south of Italy.

Every day at 10:00 or 9:00 AM, depending on legal or solar time usage, the app retrieves the birthdays from our associates "database" and sends them an e-mail full of kind wishes.
This feature was previously implemented using Google's AppScript. However, even the small maintenance was hard on the platform and I decided to migrate it to GitHub Actions, which offers a generous free tier, much more than needed by this small app.

## Knwon Limitations

### Rate Limits

As of october 2025, Gmail has an hard cap of 500 mails per day for personal accounts.
While much higher than needed for birthday emails (for our NPO, I've never seen more than 10 birthdays per day), this could be a limitation for future extensions of the app.

### Passwords

The authentication used to send emails is Google App password, which may be deprecated anytime. As of october 2025, there is currently no plan to disable it for Google Mail

### Extension

The app could be easily extended to support various mail providers. Most of them use basic HTTP security so would be easier to setup.
Other consumer mail providers, like Libero Mail, have higher or not set [Rate Limits](#rate-limits).


## Extension

It is planned to extend this app in order to support AI-generated newsletter emails based on our Instagram feed.