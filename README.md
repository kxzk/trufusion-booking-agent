<h3 align="center">TruFusion Booking Agent</h3>

> Why?

One of the instructors is super popular; therefore, in order to get into her class you have to book a week prior at midnight. Absolutely zero
chance I'm doing that, so here we are.

#### TODO

- something better for `user` + `pass`
- improve project structure
- add tests
- setup automated testing

#### Crontab

* Run Monday-Wednesday + Friday at midnight
* Hour -> 7 since UTC > PST by 7 hours

```bash
0 7 * * 1-3,5 trufusion
```

#### KT Class Schedule

> Week Of: 2022-04-11, ISO Week: 15

| Day | Time | Class | MBO ID |
| --- | --- | --- | --- |
| Mon. | 8:30 AM | Hot Pilates | 113434 |
| Mon. | 10:15 AM | Barefoot Bootcamp | 113727 |
| Tue. | 9:45 AM | Hot Pilates | 117479 |
| Tue. | 11:30 AM | Barefoot Bootcamp | 118218 |
| Wed. | 10:15 AM | Barefoot Bootcamp | 108699 |
| Fri. | 12:00 PM | Kettlebell | 108044 |

#### Class Link Info

> DATE

* Date (Generic): `Mon. Apr 4 2022 830 am`
* Date (URL): `Mon.+Apr++4%2C+2022++8%3A30+am`
* Date (Insert): `{DOW}.+{Mon}++{Day}%2C+{Year}++{Hour}%3A{Minute}+{am/pm}`

> CLASS

* Hot Pilates: `60+Min.+Tru+Hot+Pilates+%28All+Levels%29`
* Barefoot Bootcamp: `60+Min.+Tru+Barefoot+Bootcamp+%28All+Levels%29`
* Tru Kettlebell: `60+Min.+Tru+Kettlebell+%28All+Levels%29`

> CLASS ID

* each class corresponds to value on ISO WEEK 14
* each class ID increments by one per week
* get future week ISO - 14 = offset
* (iso week 14 ID value) + (offset) = (future class ID)

```python3
link = f"https://cart.mindbodyonline.com/sites/14486/cart/add_booking?item%5Binfo%5D={DATE}
&item%5Bmbo_id%5D={CLASS_ID}
&item%5Bmbo_location_id%5D=1&item%5Bname%5D={CLASS}&item%5Btype%5D=Class"
```
