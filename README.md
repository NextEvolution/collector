Collector - Facebook API Crawler
===
**Notice: this project is abandoned**

# MyGirlMel

## What is it?
MyGirlMel is a project I started in early 2016. It was designed to help sales consultants automate and better manage
their facebook sales. This is focused primarily on LuLaRoe, but could be generalized to other audiences.

## Primary Use Case
Consultants will post pictures of inventory in a facebook group photo album. Customers would comment "sold" on pictures
to buy the item.

## Components
* Capitulum - the API portion, serving up data for the front-end
* Collector - A facebook scraper - used to grab all kinds of data from facebook
* Controller - A Scheduler for facebook scraping to ensure that we stay within the call limits of the facebook API
* Data Services - There would be many data services, but these are responsible for storing data. Redis is currently
used.
* UI-Frontend - The UI for consultants

## Flow
1. A consultant logs into the service with facebook, granting access to group albums.
1. The user selects which albums to scan for sales
1. The controller places these albums on the schedule for periodic ripping
1. "sold" items are recorded
1. The consultant then views their sales in the UI

## Design Decisions
I wanted all communications to run on [NATS](https://nats.io/) - a lightweight messaging system. I chose this because
the components above are just the beginning of many other services which will require integration and I wanted to keep
communication light and fast. Capitulum would be the boundary where the HTTP API meets internal communication.

I chose [Golang](https://golang.org/) as the primary language for performance reasons. I also chose it because I believe
it is faster to write in than other systems languages. So, you get performance and somewhat reduced development time.
For services that needed to be created quickly, I would opt for writing them in Ruby and then later rewriting in Go.

---
# About Collector
The collector is designed to pull various information from the Facebook API

Specifically the following objects:
- Users Groups
- Group Albums
- Album Pictures
- Picture Comments
- Basic User Info
- Long Time Token (Allowing this app to get a longer session with facebook)
