package main_test

var postrequest = `{
  "invoice-nr": "2015-11-04T00:00:00.000Z",
  "author": "Max Mustermann",
  "city": "Musterstadt",
  "from": [
    "Musterstraße 37",
    "12345 Musterstadt"
  ],
  "to": [
    "Erika Mustermann",
    "Musterallee 1",
    "12345 Musterstadt",
    "Germany"
  ],
  "VAT": 20,
  "service": [
    {
      "description": "The first service provided",
      "price": 320,
      "details": null
    },
    {
      "description": "And another one, with a list of details",
      "price": 245,
      "details": [
        "Some more detailed explanation",
        "of the service provided",
        "Looking good"
      ]
    },
    {
      "description": "The last service provided",
      "price": 65
    }
  ],
  "closingnote": "Please transfer the due amount to the following bank account within the next 14 days:\n\n  Mustermann GmbH\n  Kreditinstitut: Deutsche Postbank AG\n  IBAN: DE18 3601 0043 9999 9999 99\n  BIC: PBNKDEFF\n\n  We really appreciate your business and look forward to future projects together.\n\n  Best regards,\n",
  "currency": "EUR",
  "lang": "english",
  "seriffont": "Hoefler Text",
  "sansfont": "Helvetica Neue",
  "fontsize": "10pt",
  "geometry": "a4paper, left=43mm, right=43mm, top=51mm, bottom=17mm"
}
`
