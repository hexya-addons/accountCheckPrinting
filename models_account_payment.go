package account_check_printing

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

//import math
func init() {
	h.AccountRegisterPayments().DeclareModel()

	h.AccountRegisterPayments().AddFields(map[string]models.FieldDefinition{
		"CheckAmountInWords": models.CharField{
			String: "Amount in Words",
		},
		"CheckManualSequencing": models.BooleanField{
			Related:  `JournalId.CheckManualSequencing`,
			ReadOnly: true,
		},
		"CheckNumber": models.IntegerField{
			String:   "Check Number",
			ReadOnly: true,
			NoCopy:   true,
			Default:  models.DefaultValue(0),
			Help: "Number of the check corresponding to this payment. If your" +
				"pre-printed check are not already numbered, you can manage" +
				"the numbering in the journal configuration page.",
		},
	})
	h.AccountRegisterPayments().Methods().OnchangeJournalId().DeclareMethod(
		`OnchangeJournalId`,
		func(rs m.AccountRegisterPaymentsSet) {
			//        if hasattr(super(AccountRegisterPayments, self), '_onchange_journal_id'):
			//            super(AccountRegisterPayments, self)._onchange_journal_id()
			//        if self.journal_id.check_manual_sequencing:
			//            self.check_number = self.journal_id.check_sequence_id.number_next_actual
		})
	h.AccountRegisterPayments().Methods().OnchangeAmount().DeclareMethod(
		`OnchangeAmount`,
		func(rs m.AccountRegisterPaymentsSet) {
			//        if hasattr(super(AccountRegisterPayments, self), '_onchange_amount'):
			//            super(AccountRegisterPayments, self)._onchange_amount()
			//        self.check_amount_in_words = self.env['account.payment']._get_check_amount_in_words(
			//            self.amount)
		})
	h.AccountRegisterPayments().Methods().GetPaymentVals().DeclareMethod(
		`GetPaymentVals`,
		func(rs m.AccountRegisterPaymentsSet) {
			//        res = super(AccountRegisterPayments, self).get_payment_vals()
			//        if self.payment_method_id == self.env.ref('account_check_printing.account_payment_method_check'):
			//            res.update({
			//                'check_amount_in_words': self.check_amount_in_words,
			//            })
			//        return res
		})
	h.AccountPayment().DeclareModel()

	h.AccountPayment().AddFields(map[string]models.FieldDefinition{
		"CheckAmountInWords": models.CharField{
			String: "Amount in Words",
		},
		"CheckManualSequencing": models.BooleanField{
			Related:  `JournalId.CheckManualSequencing`,
			ReadOnly: true,
		},
		"CheckNumber": models.IntegerField{
			String:   "Check Number",
			ReadOnly: true,
			NoCopy:   true,
			Help: "The selected journal is configured to print check numbers." +
				"If your pre-printed check paper already has numbers or" +
				"if the current numbering is wrong, you can change it in" +
				"the journal configuration page.",
		},
	})
	h.AccountPayment().Methods().GetCheckAmountInWords().DeclareMethod(
		`GetCheckAmountInWords`,
		func(rs m.AccountPaymentSet, amount interface{}) {
			//        check_amount_in_words = amount_to_text_en.amount_to_text(
			//            math.floor(amount), lang='en', currency='')
			//        check_amount_in_words = check_amount_in_words.replace(
			//            ' and Zero Cent', '')  # Ugh
			//        decimals = amount % 1
			//        if not float_is_zero(decimals, precision_digits=2):
			//            check_amount_in_words += _(' and %s/100') % str(
			//                int(round(float_round(decimals*100, precision_rounding=1))))
			//        return check_amount_in_words
		})
	h.AccountPayment().Methods().OnchangeJournalId().DeclareMethod(
		`OnchangeJournalId`,
		func(rs m.AccountPaymentSet) {
			//        if hasattr(super(AccountPayment, self), '_onchange_journal_id'):
			//            super(AccountPayment, self)._onchange_journal_id()
			//        if self.journal_id.check_manual_sequencing:
			//            self.check_number = self.journal_id.check_sequence_id.number_next_actual
		})
	h.AccountPayment().Methods().OnchangeAmount().DeclareMethod(
		`OnchangeAmount`,
		func(rs m.AccountPaymentSet) {
			//        if hasattr(super(AccountPayment, self), '_onchange_amount'):
			//            super(AccountPayment, self)._onchange_amount()
			//        self.check_amount_in_words = self._get_check_amount_in_words(
			//            self.amount)
		})
	h.AccountPayment().Methods().CheckCommunication().DeclareMethod(
		`CheckCommunication`,
		func(rs m.AccountPaymentSet, payment_method_id interface{}, communication interface{}) {
			//        super(AccountPayment, self)._check_communication(
			//            payment_method_id, communication)
			//        if payment_method_id == self.env.ref('account_check_printing.account_payment_method_check').id:
			//            if not communication:
			//                return
			//            if len(communication) > 60:
			//                raise ValidationError(
			//                    _("A check memo cannot exceed 60 characters."))
		})
	h.AccountPayment().Methods().Create().Extend(
		`Create`,
		func(rs m.AccountPaymentSet, vals models.RecordData) {
			//        if vals['payment_method_id'] == self.env.ref('account_check_printing.account_payment_method_check').id:
			//            journal = self.env['account.journal'].browse(vals['journal_id'])
			//            if journal.check_manual_sequencing:
			//                vals.update(
			//                    {'check_number': journal.check_sequence_id.next_by_id()})
			//        return super(AccountPayment, self).create(vals)
		})
	h.AccountPayment().Methods().PrintChecks().DeclareMethod(
		` Check that the recordset is valid, set the payments state
to sent and call print_checks() `,
		func(rs m.AccountPaymentSet) {
			//        self = self.filtered(lambda r: r.payment_method_id.code ==
			//                             'check_printing' and r.state != 'reconciled')
			//        if len(self) == 0:
			//            raise UserError(_("Payments to print as a checks must have 'Check' selected as payment method and "
			//                              "not have already been reconciled"))
			//        if any(payment.journal_id != self[0].journal_id for payment in self):
			//            raise UserError(
			//                _("In order to print multiple checks at once, they must belong to the same bank journal."))
			//        if not self[0].journal_id.check_manual_sequencing:
			//            # The wizard asks for the number printed on the first pre-printed check
			//            # so payments are attributed the number of the check the'll be printed on.
			//            last_printed_check = self.search([
			//                ('journal_id', '=', self[0].journal_id.id),
			//                ('check_number', '!=', 0)], order="check_number desc", limit=1)
			//            next_check_number = last_printed_check and last_printed_check.check_number + 1 or 1
			//            return {
			//                'name': _('Print Pre-numbered Checks'),
			//                'type': 'ir.actions.act_window',
			//                'res_model': 'print.prenumbered.checks',
			//                'view_type': 'form',
			//                'view_mode': 'form',
			//                'target': 'new',
			//                'context': {
			//                    'payment_ids': self.ids,
			//                    'default_next_check_number': next_check_number,
			//                }
			//            }
			//        else:
			//            self.filtered(lambda r: r.state == 'draft').post()
			//            return self.do_print_checks()
		})
	h.AccountPayment().Methods().UnmarkSent().DeclareMethod(
		`UnmarkSent`,
		func(rs m.AccountPaymentSet) {
			//        self.write({'state': 'posted'})
		})
	h.AccountPayment().Methods().DoPrintChecks().DeclareMethod(
		` This method is a hook for l10n_xx_check_printing modules
to implement actual check printing capabilities `,
		func(rs m.AccountPaymentSet) {
			//        raise UserError(_("There is no check layout configured.\nMake sure the proper check printing module is installed"
			//                          " and its configuration (in company settings > 'Configuration' tab) is correct."))
		})
}
