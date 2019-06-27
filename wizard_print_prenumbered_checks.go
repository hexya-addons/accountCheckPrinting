package account_check_printing

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.PrintPrenumberedChecks().DeclareModel()

	h.PrintPrenumberedChecks().AddFields(map[string]models.FieldDefinition{
		"NextCheckNumber": models.IntegerField{
			String:   "Next Check Number",
			Required: true,
		},
	})
	h.PrintPrenumberedChecks().Methods().PrintChecks().DeclareMethod(
		`PrintChecks`,
		func(rs m.PrintPrenumberedChecksSet) {
			//        check_number = self.next_check_number
			//        payments = self.env['account.payment'].browse(
			//            self.env.context['payment_ids'])
			//        payments.filtered(lambda r: r.state == 'draft').post()
			//        payments.filtered(lambda r: r.state != 'sent').write({'state': 'sent'})
			//        for payment in payments:
			//            payment.check_number = check_number
			//            check_number += 1
			//        return payments.do_print_checks()
		})
}
